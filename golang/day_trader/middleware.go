package main

import (
	"context"
	"errors"

	pb "./protobuff"
	"google.golang.org/grpc"
)

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	if command, ok := req.(*pb.Command); ok {
		if err := checks(command); err != nil {
			return nil, err
		}
	}
	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

// make sure user exists
func checks(req *pb.Command) error {
	createUser(req.UserId)
	if req.Name != "DUMPLOG" && req.Name != "DISPLAY_SUMMARY" {
		if req.UserId == "" {
			return errors.New("No user Id specified")
		}
	}
	return nil
}
