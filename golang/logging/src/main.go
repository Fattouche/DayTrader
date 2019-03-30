package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// This server implements the protobuff Node type
type server struct{}

var (
	grpcPort = ":40000"
)

func (s *server) LogUserCommand(ctx context.Context, req *pb.Log) (*pb.Response, error) {
	_, err := db.Exec(
		"INSERT INTO UserCommandLog(Server, TransactionNum, Command,"+
			"Username, StockSymbol, Filename, Funds) VALUES(?,?,?,?,?,?,?)",
		req.ServerName, req.TransactionNum, req.Command, req.Username,
		req.StockSymbol, req.Filename, req.Funds,
	)
	if err != nil {
		log.Println(err)
	}
	return &pb.Response{Message: "Inserted"}, err
}

func (s *server) LogQuoteServerEvent(ctx context.Context, req *pb.Log) (*pb.Response, error) {
	_, err := db.Exec(
		"INSERT INTO QuoteServerLog(Server, TransactionNum, Price,"+
			"StockSymbol, Username, QuoteServerTime, CryptoKey)"+
			"VALUES(?,?,?,?,?,?,?)",
		req.ServerName, req.TransactionNum, req.Price, req.StockSymbol,
		req.Username, req.QuoteServerTime, req.CryptoKey,
	)
	if err != nil {
		log.Println(err)
	}
	return &pb.Response{Message: "Inserted"}, err
}

func (s *server) LogAccountTransaction(ctx context.Context, req *pb.Log) (*pb.Response, error) {
	_, err := db.Exec(
		"INSERT INTO AccountTransactionLog(Server, TransactionNum, Action,"+
			"Username, Funds) VALUES(?,?,?,?,?)",
		req.ServerName, req.TransactionNum, req.AccountAction,
		req.Username, req.Funds,
	)
	if err != nil {
		log.Println(err)
	}
	return &pb.Response{Message: "Inserted"}, err
}

func (s *server) LogSystemEvent(ctx context.Context, req *pb.Log) (*pb.Response, error) {
	_, err := db.Exec(
		"INSERT INTO SystemEventLog(Server, TransactionNum, Command,"+
			"Username, StockSymbol, Filename, Funds) VALUES(?,?,?,?,?,?,?)",
		req.ServerName, req.TransactionNum, req.Command, req.Username,
		req.StockSymbol, req.Filename, req.Funds,
	)
	if err != nil {
		log.Println(err)
	}
	return &pb.Response{Message: "Inserted"}, err
}

func (s *server) LogErrorEvent(ctx context.Context, req *pb.Log) (*pb.Response, error) {
	_, err := db.Exec(
		"INSERT INTO ErrorEventLog(Server, TransactionNum, Command,"+
			"Username, StockSymbol, Filename, Funds, ErrorMessage)"+
			"VALUES(?,?,?,?,?,?,?,?)",
		req.ServerName, req.TransactionNum, req.Command, req.Username,
		req.StockSymbol, req.Filename, req.Funds, req.ErrorMessage,
	)
	if err != nil {
		log.Println(err)
	}
	return &pb.Response{Message: "Inserted"}, err
}

func (s *server) LogDebugEvent(ctx context.Context, req *pb.Log) (*pb.Response, error) {
	_, err := db.Exec(
		"INSERT INTO DebugEventLog(Server, TransactionNum, Command,"+
			"Username, StockSymbol, Filename, Funds, DebugMessage)"+
			"VALUES(?,?,?,?,?,?,?)",
		req.ServerName, req.TransactionNum, req.Command, req.Username,
		req.StockSymbol, req.Filename, req.Funds, req.DebugMessage,
	)
	if err != nil {
		log.Println(err)
	}
	return &pb.Response{Message: "Inserted"}, err
}

func (s *server) DumpLogs(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	go dumpLogsToXML(req.UserId, req.Filename)

	return &pb.Response{UserId: req.UserId, Message: "Writing to XML"}, nil
}

func (s *server) DisplaySummary(ctx context.Context, req *pb.Command) (*pb.SummaryResponse, error) {
	rows, err := getRows("UserCommandLog", req.UserId)
	if err != nil {
		return &pb.SummaryResponse{}, err
	}

	var entries []pb.Transaction

	for rows.Next() {
		entry := pb.Transaction{}
		var timestamp time.Time
		err = rows.Scan(
			&timestamp, nil, nil,
			&entry.CommandName, nil, &entry.StockSymbol,
			nil, &entry.StockAmount,
		)
		if err != nil {
			log.Println("Error scanning rows: ", err)
		}
		entry.Timestamp = strconv.FormatInt(timestamp.UnixNano()/1000000, 10)

		entries = append(entries, entry)
	}
	rows.Close()

	log.Println(entries)

	// TODO: aggregate all transactions and return
	return &pb.SummaryResponse{}, nil
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("Running logging server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	createAndOpenDB()
	startGRPCServer()
}
