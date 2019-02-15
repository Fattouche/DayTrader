package main

import (
	"context"
	"encoding/xml"
	"log"
	"net"
	"os"
	"path/filepath"

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

type UserCommandLog struct {
	XMLName        xml.Name `xml:"userCommand"`
	Timestamp      int64    `xml:"timestamp"`
	Server         string   `xml:"server"`
	TransactionNum int      `xml:"transactionNum"`
	Command        string   `xml:"command"`
	Username       string   `xml:"username"`
	StockSymbol    string   `xml:"stockSymbol"`
	Filename       string   `xml:"filename"`
	Funds          float32  `xml:"funds"`
}

func dumpLogsToXML(userID string, filename string) {
	userFilter := "*"
	if userID != "" {
		userFilter = userID
	}

	executable, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable: ", err)
	}
	currDir := filepath.Dir(executable)
	sanitizedFilename := filepath.Join(currDir, "dumplogs", filepath.Clean(filename))
	f, err := os.Create(sanitizedFilename)
	if err != nil {
		log.Println("Error scanning trigger: ", err)
	}
	defer f.Close()

	f.WriteString("<?xml version=\"1.0\"?>\n")
	f.WriteString("<log>\n")

	rows, err := db.Query("SELECT * FROM UserCommandLog WHERE Username = ?",
		userFilter)
	for rows.Next() {
		xmlLog := &UserCommandLog{}
		err = rows.Scan(
			&xmlLog.Timestamp, &xmlLog.Server, &xmlLog.TransactionNum,
			&xmlLog.Command, &xmlLog.Username, &xmlLog.StockSymbol,
			&xmlLog.Filename, &xmlLog.Funds,
		)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		output, err := xml.MarshalIndent(xmlLog, "\t", "\t")
		if err != nil {
			log.Println("Error marshalling to XML: ", err)
		}
		f.Write(output)
		f.WriteString("\n")
	}
	rows.Close()

	f.WriteString("</log>\n")
}

func (s *server) DumpLogs(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	go dumpLogsToXML(req.UserId, req.Filename)

	return &pb.Response{Message: "Writing to XML"}, nil
}

func (s *server) DumpLogs(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	return &pb.Response{Message: "yee dump"}, nil
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
