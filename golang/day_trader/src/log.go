package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"
	"google.golang.org/grpc"
)

var logUrl = "logging_lb:80"
var logChan = make(chan *logObj, 10000)

type logObj struct {
	log      *pb.Log
	funcName string
}

func makeLogFromContext(ctx context.Context) (pb.Log, error) {
	if contextValue := ctx.Value(logKey("log")); contextValue != nil {
		if log, ok := contextValue.(pb.Log); ok {
			return log, nil
		} else {
			return pb.Log{}, errors.New("Context log wasn't of log type")
		}
	}

	// This is needed because triggers can cause logs, which are run by a job
	// rather than a transaction/user command
	return pb.Log{
		TransactionNum: 1,
		Username:       "__no_user__",
		ServerName:     "Beaver_1", // TODO(cailan): use environment variable
		Command:        "__no_command__",
	}, nil
}

// Populates the context with logging info, so that it can be extracted whenever
// a log event occurs
func populateContextWithLogInfo(ctx context.Context, command *pb.Command) context.Context {
	var key logKey

	// TODO: get the server name from an environment variable
	rawLog := pb.Log{
		TransactionNum: command.TransactionId,
		Username:       command.UserId,
		ServerName:     "Beaver_1",
		Command:        command.Name,
	}
	key = "log"
	ctx = context.WithValue(ctx, key, rawLog)
	return ctx
}

func logUserCommand(ctx context.Context, command *pb.Command) error {
	logObject, err := makeLogFromContext(ctx)
	if err != nil {
		log.Println("Error making log from context: ", err)
		return err
	}
	logObject.StockSymbol = command.Symbol
	logObject.Filename = command.Filename
	logObject.Funds = command.Amount
	logEvent := &logObj{log: &logObject, funcName: "LogUserCommand"}
	logChan <- logEvent
	return nil
}

func logErrorEvent(ctx context.Context, logErr error) error {
	logObject, err := makeLogFromContext(ctx)
	if err != nil {
		log.Println("Error making log from context: ", err)
		return err
	}
	logObject.ErrorMessage = logErr.Error()
	logEvent := &logObj{log: &logObject, funcName: "LogErrorEvent"}
	logChan <- logEvent
	return nil
}

func startLoggerWorker() {
	conn, err := grpc.Dial(logUrl, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial to %s with %v", logUrl, err)
	}
	client := pb.NewLoggerClient(conn)
	for {
		obj := <-logChan
		switch obj.funcName {
		case "LogUserCommand":
			client.LogUserCommand(context.Background(), obj.log)
		case "LogQuoteServerEvent":
			client.LogQuoteServerEvent(context.Background(), obj.log)
		case "LogAccountTransaction":
			client.LogAccountTransaction(context.Background(), obj.log)
		case "LogSystemEvent":
			client.LogSystemEvent(context.Background(), obj.log)
		case "LogErrorEvent":
			client.LogErrorEvent(context.Background(), obj.log)
		case "LogDebugEvent":
			client.LogDebugEvent(context.Background(), obj.log)
		}
	}
}
