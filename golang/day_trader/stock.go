package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Stock struct {
	Symbol    string
	Price     float32
	Hash      string
	TimeStamp time.Time
}

func quote(userID, symbol string) (*Stock, error) {
	stock, err := c.getCacheStock(symbol)
	if err != nil || (stock != nil && stock.isExpired()) {
		//TODO: Do something with this hash
		stock = &Stock{Symbol: symbol}
		stock.Price, stock.Hash, err = executeRequest(userID, symbol)
		stock.TimeStamp = time.Now()
		c.setCache(stock.Symbol, stock)
		return stock, err
	}
	return stock, err
}

func (stock *Stock) isExpired() bool {
	duration := time.Since(stock.TimeStamp)
	if duration > time.Second*60 {
		return true
	}
	return false
}

func executeRequest(userID, symbol string) (float32, string, error) {
	ln, err := net.Dial("tcp", QUOTE_HOST+QUOTE_PORT)
	defer ln.Close()
	if err != nil {
		return 0, "", err
	}
	buf := make([]byte, 300)
	str := fmt.Sprintf("%s,%s\r", symbol, userID)
	ln.Write([]byte(str))
	len, err := ln.Read(buf)
	if err != nil {
		return 0, "", err
	}
	info := string(buf[:len])
	infoArr := strings.Split(info, ",")
	price, err := strconv.ParseFloat(infoArr[0], 32)
	if err != nil {
		return 0, "", err
	}
	return float32(price), infoArr[4], nil
}
