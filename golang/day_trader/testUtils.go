package main

import (
	"encoding/json"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"
)

var s server
var symbol = "ABC"
var userId = "tester"
var filename = "dumplog"
var amount = float32(521)
var transactionId = int32(1)
var quotePrice = float32(5.21)
var hash = "lod23EP0lofFCkEd0ilcUpjL0MuBcIh3HiwAq9QSXdU="
var name = "tester"

func genGrpcRequest(name string) *pb.Command {
	req := &pb.Command{UserId: userId, Symbol: symbol, Amount: amount, TransactionId: transactionId, Name: name, Filename: filename}
	return req
}

func toString(msg interface{}) string {
	bytes, _ := json.Marshal(msg)
	return string(bytes)
}
