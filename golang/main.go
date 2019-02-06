package main

import (
	"context"
	"fmt"
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

func (s *server) Add(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user, err := getUser(req.UserId)
	if err != nil {
		return nil, err
	}
	user.updateUserBalance(req.Amount)
	msg := fmt.Sprintf("New balance %f", user.Balance)
	return &pb.Response{Message: msg}, nil
}

func (s *server) Buy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) Quote(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	stock, err := quote(req.UserId, req.Symbol)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("Stock: %s, price: %f", stock.Symbol, stock.Price)
	return &pb.Response{Message: msg}, nil
}

func (s *server) Sell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) CommitBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}
func (s *server) CommitSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) CancelBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
}

func (s *server) CancelSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee"}, nil
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
