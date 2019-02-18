package main

import (
	pb "github.com/Fattouche/DayTrader/golang/protobuff"
)

var s server
var symbol = "ABC"
var testUserId = "tester"
var filename = "dumplog"
var amount = float32(521)
var transactionId = int32(1)
var quotePrice = float32(5.21)
var hash = "lod23EP0lofFCkEd0ilcUpjL0MuBcIh3HiwAq9QSXdU="
var testName = "tester"

func genGrpcRequest(name string) *pb.Command {
	req := &pb.Command{UserId: testUserId, Symbol: symbol, Amount: amount, TransactionId: transactionId, Name: testName, Filename: filename}
	return req
}
