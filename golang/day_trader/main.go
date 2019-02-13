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
	Name      string
	Id        string
	BuyStack  []*Buy
	SellStack []*Sell
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
	GRPC_PORT  = ":41000"
	DB_NAME    = "daytrader"
	QUOTE_HOST = "quote_server"
	QUOTE_PORT = ":4442"
	CACHE_HOST = "cache"
	CACHE_PORT = ":11211"
)

func (s *server) Add(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	user.updateUserBalance(req.Amount)
	return &pb.Response{Message: user.toString()}, nil
}

func (s *server) Buy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	buy, err := createBuy(req.Amount, req.Symbol, user.Id)
	return &pb.Response{Message: buy.toString()}, err
}

func (s *server) Quote(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	stock, err := quote(req.UserId, req.Symbol)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Message: stock.toString()}, nil
}

func (s *server) Sell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	sell, err := createSell(req.Amount, req.Symbol, user.Id)
	if err != nil {
		return nil, err
	}
	user.SellStack = append(user.SellStack, sell)
	user.setCache()
	return &pb.Response{Message: sell.toString()}, nil
}

func (s *server) CommitBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	buy := user.popFromBuyStack()
	if buy == nil {
		return nil, errors.New("No buy on the stack")
	}
	userStock, err := buy.commit(false)
	return &pb.Response{Message: userStock.toString()}, err
}
func (s *server) CommitSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	sell := user.popFromSellStack()
	if sell == nil {
		return nil, errors.New("No sell on the stack")
	}
	err := sell.commit(false)
	return &pb.Response{Message: user.toString()}, err
}

func (s *server) CancelBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	buy := user.popFromBuyStack()
	if buy != nil {
		buy.cancel()
	} else {
		return nil, errors.New("No buy on stack")
	}
	return &pb.Response{Message: user.toString()}, nil
}

func (s *server) CancelSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	sell := user.popFromSellStack()
	if sell != nil {
		sell.cancel()
	} else {
		return nil, errors.New("No sell on stack")
	}
	return &pb.Response{Message: user.toString()}, nil
}

func (s *server) SetBuyAmount(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	if user.Balance < req.Amount {
		return nil, fmt.Errorf("Not enough balance, have %f need %f", user.Balance, req.Amount)
	}
	trigger, err := getBuyTrigger(user.Id, req.Symbol)
	if err != nil {
		buy, err := createBuy(req.Amount, req.Symbol, user.Id)
		if err != nil {
			return nil, err
		}
		buy, err = buy.insertBuy()
		if err != nil {
			log.Println(err)
		}
		trigger := createBuyTrigger(user.Id, req.Symbol, buy.Id, req.Amount)
		return &pb.Response{Message: trigger.toString()}, nil
	}
	return &pb.Response{Message: trigger.toString()}, trigger.updateCashAmount(req.Amount)
}

func (s *server) SetSellAmount(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getSellTrigger(req.UserId, req.Symbol)
	if err != nil {
		sell := &Sell{StockSymbol: req.Symbol, UserId: req.UserId}
		err = sell.updateCashAmount(req.Amount)
		if err != nil {
			return nil, err
		}
		sell.insertSell()
		trigger := createSellTrigger(req.UserId, req.Symbol, sell.Id, req.Amount)
		return &pb.Response{Message: trigger.toString()}, nil
	}
	return &pb.Response{Message: trigger.toString()}, trigger.updateCashAmount(req.Amount)
}

func (s *server) SetBuyTrigger(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getBuyTrigger(req.UserId, req.Symbol)
	if err != nil {
		return nil, errors.New("Trigger requires a buy amount first, please make one")
	}
	trigger.updatePrice(req.Amount)
	return &pb.Response{Message: trigger.toString()}, nil
}

func (s *server) SetSellTrigger(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getSellTrigger(req.UserId, req.Symbol)
	if err != nil {
		return nil, errors.New("Trigger requires a sell amount first, please make one")
	}
	trigger.updatePrice(req.Amount)
	return &pb.Response{Message: trigger.toString()}, nil
}

func (s *server) CancelSetBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getBuyTrigger(req.UserId, req.Symbol)
	if err != nil {
		return nil, errors.New("Set buy not found")
	}
	trigger.cancel()
	return &pb.Response{Message: "Disabling trigger"}, nil
}

func (s *server) CancelSetSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getSellTrigger(req.UserId, req.Symbol)
	if err != nil {
		return nil, errors.New("Set sell not found")
	}
	trigger.cancel()
	return &pb.Response{Message: "Disabling trigger"}, nil
}

func (s *server) DumpLog(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) DisplaySummary(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
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
	startGRPCServer()
	go watchTriggers()
}
