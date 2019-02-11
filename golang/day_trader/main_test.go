package main

import (
	"context"
	"testing"
	"time"
)

func init() {
	c = &mockCache{}
	db = &mockDB{}
}

func TestQuote(t *testing.T) {
	s := &server{}
	resp, err := s.Quote(context.Background(), genGrpcRequest("QUOTE"))
	exp := &Stock{Price: float32(quotePrice), Symbol: symbol, Hash: hash, TimeStamp: time.Now()}
	if err != nil {
		t.Error("TestQuote got unexpected error: ", err)
	}
	if resp.Message != toString(exp) {
		t.Errorf("TestQuote got %v wanted %v", resp.Message, toString(exp))
	}
}

func TestAdd(t *testing.T) {
	s := &server{}
	resp, err := s.Add(context.Background(), genGrpcRequest("ADD"))
	exp := &User{Balance: amount, Id: userId, Name: name, BuyStack: []*Buy{}, SellStack: []*Sell{}}
	if err != nil {
		t.Error("TestAdd got unexpected error: ", err)
	}
	if resp.Message != toString(exp) {
		t.Errorf("TestAdd got %v wanted %v", resp.Message, toString(exp))
	}
}
