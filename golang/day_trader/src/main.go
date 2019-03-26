package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	_ "net/http/pprof"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//easyjson:json
type Stock struct {
	Symbol    string
	Price     float32
	Hash      string
	TimeStamp time.Time
}

//easyjson:json
type User struct {
	Balance   float32
	Password  string
	Id        string
	BuyStack  []*Buy
	SellStack []*Sell
	StockMap  map[string]int32
}

//easyjson:json
type Buy struct {
	Id                 int64
	Price              float32
	StockSymbol        string
	IntendedCashAmount float32
	ActualCashAmount   float32
	StockBoughtAmount  int
	UserId             string
	Timestamp          time.Time
}

//easyjson:json
type Sell struct {
	Id                 int64
	Price              float32
	StockSymbol        string
	IntendedCashAmount float32
	ActualCashAmount   float32
	StockSoldAmount    int
	UserId             string
	Timestamp          time.Time
}

//easyjson:json
type BuyTrigger struct {
	UserId string
	BuyId  int64
	Active bool
}

//easyjson:json
type SellTrigger struct {
	UserId string
	SellId int64
	Active bool
}

//easyjson:json
type UserStock struct {
	UserId      string
	StockSymbol string
	Amount      int
}

// This server implements the protobuff Node type
type server struct{}

var (
	GRPC_PORT   = ":41000"
	DB_NAME     = "daytrader"
	QUOTE_HOST  = "quote_server"
	QUOTE_PORT  = ":4442"
	CACHE_HOST  = "cache"
	CACHE_PORT  = ":11211"
	MAX_BALANCE = float32(1000000000000)
)

func (s *server) GetUser(ctx context.Context, req *pb.Command) (*pb.UserResponse, error) {
	balance, stocks, err := userExists(req.UserId, req.Password)
	var userResp *pb.UserResponse
	if err == nil {
		userResp = &pb.UserResponse{
			UserId:  req.UserId,
			Balance: balance,
			Stocks:  stocks,
		}
	}
	return userResp, err
}

func (s *server) CreateUser(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{UserId: req.UserId, Message: "Successfully created user"}, createUser(req.UserId, req.Password)
}

func (s *server) Add(ctx context.Context, req *pb.Command) (*pb.BalanceResponse, error) {
	user := getUser(req.UserId)
	user, err := user.updateUserBalance(ctx, req.Amount, true)
	if err == nil {
		err = errors.New(fmt.Sprintf("Amount: %f Balance: %f", req.Amount, user.Balance))
	}
	return &pb.BalanceResponse{UserId: user.Id, Balance: user.Balance}, err
}

func (s *server) Buy(ctx context.Context, req *pb.Command) (*pb.BalanceResponse, error) {
	user := getUser(req.UserId)
	_, err := createBuy(ctx, req.Amount, req.Symbol, user)
	if err != nil {
		logErrorEvent(ctx, err)
	}
	return &pb.BalanceResponse{UserId: user.Id, Balance: user.Balance}, err
}

func (s *server) Quote(ctx context.Context, req *pb.Command) (*pb.PriceResponse, error) {
	stock, err := quote(ctx, req.UserId, req.Symbol)
	if err != nil {
		logErrorEvent(ctx, err)
		return nil, err
	}

	return &pb.PriceResponse{UserId: req.UserId, Price: stock.Price}, nil
}

func (s *server) Sell(ctx context.Context, req *pb.Command) (*pb.StockUpdateResponse, error) {
	user := getUser(req.UserId)
	sell, err := createSell(ctx, req.Amount, req.Symbol, user)
	if err != nil {
		logErrorEvent(ctx, err)
		return nil, err
	}
	user.SellStack = append(user.SellStack, sell)
	user.setCache()
	return &pb.StockUpdateResponse{UserId: user.Id, Stocks: map[string]int32{req.Symbol: user.StockMap[req.Symbol]}}, nil
}

func (s *server) CommitBuy(ctx context.Context, req *pb.Command) (*pb.StockUpdateResponse, error) {
	var err error
	user := getUser(req.UserId)
	buy := user.popFromBuyStack()
	if buy == nil {
		err = errors.New("No buy on the stack")
		logErrorEvent(ctx, err)
		return nil, err
	}
	userStock, err := buy.commit(ctx, user, false)
	if err != nil {
		logErrorEvent(ctx, err)
		return nil, err
	}
	return &pb.StockUpdateResponse{UserId: user.Id, Stocks: map[string]int32{userStock.StockSymbol: int32(userStock.Amount)}}, err
}

func (s *server) CommitSell(ctx context.Context, req *pb.Command) (*pb.UserResponse, error) {
	var err error
	user := getUser(req.UserId)
	sell := user.popFromSellStack()
	if sell == nil {
		err = errors.New("No sell on the stack")
		logErrorEvent(ctx, err)
		return nil, err
	}
	err = sell.commit(ctx, false, user)
	if err != nil {
		logErrorEvent(ctx, err)
	}
	return &pb.UserResponse{UserId: user.Id, Balance: user.Balance, Stocks: map[string]int32{sell.StockSymbol: user.StockMap[sell.StockSymbol]}}, err
}

func (s *server) CancelBuy(ctx context.Context, req *pb.Command) (*pb.BalanceResponse, error) {
	user := getUser(req.UserId)
	buy := user.popFromBuyStack()
	if buy != nil {
		buy.cancel(ctx, user)
	} else {
		err := errors.New("No buy on stack")
		logErrorEvent(ctx, err)
		return nil, err
	}
	return &pb.BalanceResponse{UserId: user.Id, Balance: user.Balance}, nil
}

func (s *server) CancelSell(ctx context.Context, req *pb.Command) (*pb.StockUpdateResponse, error) {
	user := getUser(req.UserId)
	sell := user.popFromSellStack()
	if sell != nil {
		sell.cancel(ctx, user)
	} else {
		err := errors.New("No sell on stack")
		logErrorEvent(ctx, err)
		return nil, err
	}
	return &pb.StockUpdateResponse{UserId: user.Id, Stocks: map[string]int32{sell.StockSymbol: user.StockMap[sell.StockSymbol]}}, nil
}

func (s *server) SetBuyAmount(ctx context.Context, req *pb.Command) (*pb.BalanceResponse, error) {
	var err error
	user := getUser(req.UserId)
	if user.Balance < req.Amount {
		err = fmt.Errorf("Not enough balance, have %f need %f", user.Balance, req.Amount)
		logErrorEvent(ctx, err)
		return nil, err
	}
	trigger, err := getBuyTrigger(ctx, user.Id, req.Symbol)
	if err != nil {
		buy, err := createBuy(ctx, req.Amount, req.Symbol, user)
		if err != nil {
			return nil, err
		}
		buy, err = buy.insertBuy(ctx)
		if err != nil {
			logErrorEvent(ctx, err)
			log.Println(err)
		}
		createBuyTrigger(ctx, user.Id, req.Symbol, buy.Id, req.Amount)
		return &pb.BalanceResponse{UserId: user.Id, Balance: user.Balance}, nil
	}
	err = trigger.updateCashAmount(ctx, req.Amount, user)
	if err != nil {
		logErrorEvent(ctx, err)
	}
	return &pb.BalanceResponse{UserId: user.Id, Balance: user.Balance}, err
}

func (s *server) SetSellAmount(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	trigger, err := getSellTrigger(ctx, req.UserId, req.Symbol)
	if err != nil {
		user := getUser(req.UserId)
		sell := &Sell{StockSymbol: req.Symbol, UserId: req.UserId}
		err = sell.updateCashAmount(ctx, req.Amount, user)
		if err != nil {
			logErrorEvent(ctx, err)
			return nil, err
		}
		sell.insertSell(ctx)
		createSellTrigger(ctx, req.UserId, req.Symbol, sell.Id, req.Amount)
		return &pb.Response{UserId: user.Id, Message: "Successfully set sell amount"}, nil
	}
	err = trigger.updateCashAmount(ctx, req.Amount, user)
	if err != nil {
		logErrorEvent(ctx, err)
	}
	return &pb.Response{UserId: user.Id, Message: "Successfully set sell amount"}, err

}

func (s *server) SetBuyTrigger(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getBuyTrigger(ctx, req.UserId, req.Symbol)
	if err != nil {
		err = errors.New("Trigger requires a buy amount first, please make one")
		logErrorEvent(ctx, err)
		return nil, err
	}
	trigger.updatePrice(ctx, req.Amount)
	return &pb.Response{UserId: req.UserId, Message: "Successfully set buy trigger"}, nil
}

func (s *server) SetSellTrigger(ctx context.Context, req *pb.Command) (*pb.StockUpdateResponse, error) {
	user := getUser(req.UserId)
	trigger, err := getSellTrigger(ctx, req.UserId, req.Symbol)
	if err != nil {
		err = errors.New("Trigger requires a sell amount first, please make one")
		logErrorEvent(ctx, err)
		return nil, err
	}
	trigger.updatePrice(ctx, req.Amount, user)
	return &pb.StockUpdateResponse{UserId: user.Id, Stocks: map[string]int32{req.Symbol: user.StockMap[req.Symbol]}}, nil
}

func (s *server) CancelSetBuy(ctx context.Context, req *pb.Command) (*pb.BalanceResponse, error) {
	user := getUser(req.UserId)
	trigger, err := getBuyTrigger(ctx, req.UserId, req.Symbol)
	if err != nil {
		err = errors.New("Set buy not found")
		logErrorEvent(ctx, err)
		return nil, err
	}
	trigger.cancel(ctx, user)
	return &pb.BalanceResponse{UserId: user.Id, Balance: user.Balance}, nil
}

func (s *server) CancelSetSell(ctx context.Context, req *pb.Command) (*pb.StockUpdateResponse, error) {
	user := getUser(req.UserId)
	trigger, err := getSellTrigger(ctx, req.UserId, req.Symbol)
	if err != nil {
		err = errors.New("Set sell not found")
		logErrorEvent(ctx, err)
		return nil, err
	}
	trigger.cancel(ctx, user)
	return &pb.StockUpdateResponse{UserId: user.Id, Stocks: map[string]int32{req.Symbol: user.StockMap[req.Symbol]}}, nil
}

func (s *server) DumpLog(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	conn, err := grpc.Dial(logUrl, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial to %s with %v", logUrl, err)
	}
	client := pb.NewLoggerClient(conn)
	response, err := client.DumpLogs(ctx, req)
	if err != nil {
		log.Println("Dumplog failed: ", err)
	}
	return response, err
}

func (s *server) DisplaySummary(ctx context.Context, req *pb.Command) (*pb.SummaryResponse, error) {
	// conn, err := grpc.Dial(logUrl, grpc.WithInsecure())
	// if err != nil {
	// 	log.Printf("Failed to dial to %s with %v", logUrl, err)
	// }
	// client := pb.NewLoggerClient(conn)
	// _, err = client.DisplaySummary(ctx, req)
	// if err != nil {
	// 	log.Println("DisplaySummary call to log server failed: ", err)
	// }

	// _ (rename to response when ready to implement) should contain all transactions, once that's implemented
	// TODO: grab info on current triggers, and then put it all together into
	// a SummaryResponse proto and return

	return &pb.SummaryResponse{}, nil
}

// Starts a generic GRPC server
func startGRPCServer() {
	lis, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(withServerUnaryInterceptor())
	pb.RegisterDayTraderServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func watchTriggers() {
	for {
		time.Sleep(time.Second * 60)
		checkSellTriggers()
		checkBuyTriggers()
	}
}

func main() {
	//Uncomment and run `go tool pprof -png http://localhost:6060/debug/pprof/profile?seconds=30 > out.png` to get image
	// go func() {
	// 	log.Println(http.ListenAndServe(":6060", nil))
	// }()
	createAndOpenDB()
	initCache()
	for i := 0; i < 500; i++ {
		go startLoggerWorker()
	}
	go watchTriggers()
	startGRPCServer()
}
