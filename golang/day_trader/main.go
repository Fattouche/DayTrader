package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"

	pb "./protobuff"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// This server implements the protobuff Node type
type server struct{}

const (
	GRPC_PORT  = ":41000"
	DB_NAME    = "daytrader"
	QUOTE_HOST = "localhost:"
	QUOTE_PORT = "4442"
)

func toString(msg interface{}) string {
	bytes, _ := json.Marshal(msg)
	return string(bytes)
}

func (user *User) popFromBuyStack() *Buy {
	buy := user.BuyStack[len(user.BuyStack)-1]
	user.BuyStack = user.BuyStack[:len(user.BuyStack)-1]
	setCache(user.Id, user)
	return buy
}

func (user *User) popFromSellStack() *Sell {
	sell := user.SellStack[len(user.SellStack)-1]
	user.SellStack = user.SellStack[:len(user.SellStack)-1]
	setCache(user.Id, user)
	return sell
}

func (s *server) Add(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	user.updateUserBalance(req.Amount)
	return &pb.Response{Message: toString(user)}, nil
}

func (s *server) Buy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	quote, err := quote(user.Id, req.Symbol)
	if err != nil {
		return nil, err
	}
	buy, err := createBuy(quote.Price, req.Amount, float32(0), 0, req.Symbol, user.Id)
	if err != nil {
		return nil, err
	}
	user.BuyStack = append(user.BuyStack, buy)
	setCache(user.Id, user)
	return &pb.Response{Message: toString(buy)}, nil
}

func (s *server) Quote(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	stock, err := quote(req.UserId, req.Symbol)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Message: toString(stock)}, nil
}

func (s *server) Sell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	quote, err := quote(user.Id, req.Symbol)
	if err != nil {
		return nil, err
	}
	sell, err := createSell(quote.Price, req.Amount, float32(0), 0, req.Symbol, user.Id)
	if err != nil {
		return nil, err
	}
	user.SellStack = append(user.SellStack, sell)
	setCache(user.Id, user)
	return &pb.Response{Message: toString(sell)}, nil
}

func (s *server) CommitBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	if len(user.BuyStack) == 0 {
		return nil, errors.New("No buy on the stack")
	}
	buy := user.popFromBuyStack()
	userStock, err := buy.commit()
	return &pb.Response{Message: toString(userStock)}, err
}
func (s *server) CommitSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	if len(user.BuyStack) == 0 {
		return nil, errors.New("No sell on the stack")
	}
	sell := user.popFromSellStack()
	err := sell.commit()
	return &pb.Response{Message: toString(user)}, err
}

func (s *server) CancelBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	buy := user.popFromBuyStack()
	if buy != nil {
		buy.cancel()
	}
	return &pb.Response{Message: toString(user)}, nil
}

func (s *server) CancelSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	sell := user.popFromBuyStack()
	if sell != nil {
		sell.cancel()
	}
	return &pb.Response{Message: toString(user)}, nil
}

func (s *server) SetBuyAmount(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) SetSellAmount(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) SetBuyTrigger(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) SetSellTrigger(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) CancelSetSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) CancelSetBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
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

func main() {
	createAndOpenDB()
	initCache()
	startGRPCServer()
}
