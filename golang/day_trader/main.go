package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	pb "./protobuff"

	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// This server implements the protobuff Node type
type server struct{}

type cacheInterface interface {
	setCache(key string, val interface{}) error
	getCacheStock(key string) (*Stock, error)
	getCacheUser(key string) (*User, error)
}

type dbInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
}

var c cacheInterface
var db dbInterface

var (
	GRPC_PORT  = ":41000"
	DB_NAME    = "daytrader"
	QUOTE_HOST = "quote_server"
	QUOTE_PORT = ":4442"
	CACHE_HOST = "cache"
	CACHE_PORT = ":11211"
)

func toString(msg interface{}) string {
	bytes, _ := json.Marshal(msg)
	return string(bytes)
}

func (s *server) Add(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	user.updateUserBalance(req.Amount)
	return &pb.Response{Message: toString(user)}, nil
}

func (s *server) Buy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	buy, err := createBuy(req.Amount, req.Symbol, user.Id)
	if err != nil {
		return nil, err
	}

	user.BuyStack = append(user.BuyStack, buy)
	c.setCache(user.Id, user)
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
	sell, err := createSell(req.Amount, req.Symbol, user.Id)
	if err != nil {
		return nil, err
	}
	user.SellStack = append(user.SellStack, sell)
	c.setCache(user.Id, user)
	return &pb.Response{Message: toString(sell)}, nil
}

func (s *server) CommitBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	buy := user.popFromBuyStack()
	if buy == nil {
		return nil, errors.New("No buy on the stack")
	}
	userStock, err := buy.commit()
	return &pb.Response{Message: toString(userStock)}, err
}
func (s *server) CommitSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	sell := user.popFromSellStack()
	if sell == nil {
		return nil, errors.New("No sell on the stack")
	}
	err := sell.commit()
	return &pb.Response{Message: toString(user)}, err
}

func (s *server) CancelBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	buy := user.popFromBuyStack()
	if buy != nil {
		buy.cancel()
	} else {
		return nil, errors.New("No buy on stack")
	}
	return &pb.Response{Message: toString(user)}, nil
}

func (s *server) CancelSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	user := getUser(req.UserId)
	sell := user.popFromSellStack()
	if sell != nil {
		sell.cancel()
	} else {
		return nil, errors.New("No sell on stack")
	}
	return &pb.Response{Message: toString(user)}, nil
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
		return &pb.Response{Message: toString(trigger)}, nil
	}
	return &pb.Response{Message: toString(trigger)}, trigger.updateCashAmount(req.Amount)
}

func (s *server) SetSellAmount(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getSellTrigger(req.UserId, req.Symbol)
	if err != nil {
		sell, err := createSell(req.Amount, req.Symbol, req.UserId)
		if err != nil {
			return nil, err
		}
		sell.insertSell()
		trigger := createSellTrigger(req.UserId, req.Symbol, sell.Id, req.Amount)
		return &pb.Response{Message: toString(trigger)}, nil
	}
	return &pb.Response{Message: toString(trigger)}, trigger.updateCashAmount(req.Amount)
}

func (s *server) SetBuyTrigger(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getBuyTrigger(req.UserId, req.Symbol)
	if err != nil || !trigger.Active {
		return nil, errors.New("Trigger requires a buy amount first, please make one")
	}
	trigger.updatePrice(req.Amount)
	return &pb.Response{Message: toString(trigger)}, nil
}

func (s *server) SetSellTrigger(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getSellTrigger(req.UserId, req.Symbol)
	if err != nil || !trigger.Active {
		return nil, errors.New("Trigger requires a sell amount first, please make one")
	}
	trigger.updatePrice(req.Amount)
	return &pb.Response{Message: toString(trigger)}, nil
}

func (s *server) CancelSetBuy(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getBuyTrigger(req.UserId, req.Symbol)
	if err != nil {
		return nil, errors.New("Set buy not found")
	}
	if !trigger.Active {
		return nil, fmt.Errorf("No active trigger found for set_buy corresponding to %s", req.Symbol)
	}
	trigger.cancel()
	return &pb.Response{Message: "Disabling trigger"}, nil
}

func (s *server) CancelSetSell(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	trigger, err := getSellTrigger(req.UserId, req.Symbol)
	if err != nil {
		return nil, errors.New("Set sell not found")
	}
	if !trigger.Active {
		return nil, fmt.Errorf("No active trigger found for set_sell corresponding to %s", req.Symbol)
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
	createAndOpenDB()
	c = &cache{Client: memcache.New(CACHE_HOST + CACHE_PORT)}
	startGRPCServer()
	go watchTriggers()
}
