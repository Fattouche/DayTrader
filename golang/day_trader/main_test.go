package main

import (
	"context"
	"strings"
	"testing"
	"time"
)

var user *User
var stock *Stock

func init() {
	c = &mockCache{}
	db = &mockDB{}
	user = &User{Balance: amount, Id: userId, Name: name, BuyStack: []*Buy{}, SellStack: []*Sell{}}
	c.setCache(user.Id, user)
	stock = &Stock{Symbol: symbol, Price: quotePrice, Hash: hash, TimeStamp: time.Now()}
	c.setCache(stock.Symbol, stock)
}

func TestQuote(t *testing.T) {
	resp, err := s.Quote(context.Background(), genGrpcRequest("QUOTE"))
	if err != nil {
		t.Error("TestQuote got unexpected error: ", err)
	}
	if !strings.Contains(resp.Message, hash) {
		t.Errorf("TestQuote %v didnt contain %v", resp.Message, hash)
	}
}

func TestAdd(t *testing.T) {
	resp, err := s.Add(context.Background(), genGrpcRequest("ADD"))
	exp := &User{Balance: amount, Id: userId, Name: name, BuyStack: []*Buy{}, SellStack: []*Sell{}}
	if err != nil {
		t.Error("TestAdd got unexpected error: ", err)
	}
	if resp.Message != toString(exp) {
		t.Errorf("TestAdd got %v wanted %v", resp.Message, toString(exp))
	}
}

func TestBuy(t *testing.T) {
	s.Buy(context.Background(), genGrpcRequest("BUY"))
	user, _ = c.getCacheUser(userId)
	if len(user.BuyStack) == 0 {
		t.Errorf("TestBuy expected Buy stack to have length of %d but had %d", 1, 0)
	}
	buy := user.popFromBuyStack()
	if buy.Price != quotePrice {
		t.Errorf("TestBuy expected quote price of %v got %v", quotePrice, buy.Price)
	}
	if buy.StockSymbol != symbol {
		t.Errorf("TestBuy expected symbol of %v got %v", symbol, buy.StockSymbol)
	}
	if buy.StockBoughtAmount != int(amount/quotePrice) {
		t.Errorf("TestBuy expected stockBought amount of %v got %v", amount/quotePrice, buy.StockBoughtAmount)
	}
	if buy.ActualCashAmount != float32(amount/quotePrice)*quotePrice {
		t.Errorf("TestBuy expected actual cash amount amount of %v got %v", float32(amount/quotePrice)*quotePrice, buy.ActualCashAmount)
	}
}

func TestSell(t *testing.T) {
	s.Sell(context.Background(), genGrpcRequest("SELL"))
	user, _ = c.getCacheUser(userId)
	if len(user.SellStack) == 0 {
		t.Errorf("TestSell expected Sell stack to have length of %d but had %d", 1, 0)
	}
	sell := user.popFromSellStack()
	if sell.Price != quotePrice {
		t.Errorf("TestSell expected quote price of %v got %v", quotePrice, sell.Price)
	}
	if sell.StockSymbol != symbol {
		t.Errorf("TestSell expected symbol of %v got %v", symbol, sell.StockSymbol)
	}
	if sell.StockSoldAmount != int(amount/quotePrice) {
		t.Errorf("TestSell expected stockSold amount of %v got %v", amount/quotePrice, sell.StockSoldAmount)
	}
	if sell.ActualCashAmount != float32(amount/quotePrice)*quotePrice {
		t.Errorf("TestSell expected actual cash amount amount of %v got %v", float32(amount/quotePrice)*quotePrice, sell.ActualCashAmount)
	}
}
