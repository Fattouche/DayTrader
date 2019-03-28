package main

import (
	"context"
	"errors"
	"log"
	"unicode"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"
	"google.golang.org/grpc"
)

type logKey string

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func serverInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	if command, ok := req.(*pb.Command); ok {
		if command.Name != "CREATE_USER" {
			ctx = populateContextWithLogInfo(ctx, command)
			err := logUserCommand(ctx, command)
			if err != nil {
				log.Println("Error logging: ", err)
				return nil, err
			}
			if err := checks(command); err != nil {
				err := logErrorEvent(ctx, err)
				return nil, err
			}
		}
	}
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

// make sure user exists
func checks(req *pb.Command) error {
	if req.Name != "DUMPLOG" && req.Name != "DISPLAY_SUMMARY" {
		if req.UserId == "" {
			return errors.New("No user Id specified")
		}
	}
	if req.Amount < 0 {
		return errors.New("Amount must be greater than 0")
	}
	if len(req.Symbol) > 3 {
		return errors.New("Symbol must be 3 characters or less")
	}
	if !isLetter(req.Symbol) {
		return errors.New("Symbol must consist of only letters")
	}
	return nil
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
