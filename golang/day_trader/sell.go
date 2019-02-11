package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"
)

type Sell struct {
	Id                 int64
	Price              float32
	StockSymbol        string
	IntendedCashAmount float32
	ActualCashAmount   float32
	StockSoldAmount    int
	UserId             string
	Timestamp          time.Time
}

func createSell(intendedCashAmount float32, symbol, userID string) (*Sell, error) {
	stock, err := quote(userID, symbol)
	if err != nil {
		return nil, err
	}
	sell := &Sell{Price: stock.Price, StockSymbol: symbol, UserId: userID}
	err = sell.updateCashAmount(intendedCashAmount)
	if err != nil {
		return nil, err
	}
	err = sell.updatePrice(stock.Price)
	if err != nil {
		return nil, err
	}
	return sell, err
}

func (sell *Sell) updateCashAmount(amount float32) error {
	stock, _ := quote(sell.UserId, sell.StockSymbol)
	userStock := getOrCreateUserStock(sell.UserId, sell.StockSymbol)
	stockSoldAmount := int(math.Floor(float64(amount / stock.Price)))
	if stockSoldAmount > userStock.Amount {
		return fmt.Errorf("Not enough stock, have %d need %d", userStock.Amount, stockSoldAmount)
	}
	sell.IntendedCashAmount = amount
	return nil
}

func (sell *Sell) updatePrice(stockPrice float32) error {
	sell.cancel()
	userStock := getOrCreateUserStock(sell.UserId, sell.StockSymbol)
	sell.StockSoldAmount = int(math.Min(math.Floor(float64(sell.IntendedCashAmount/stockPrice)), float64(userStock.Amount)))
	if sell.StockSoldAmount <= 0 {
		return errors.New("Update trigger price failed")
	}
	sell.ActualCashAmount = float32(sell.StockSoldAmount) * stockPrice
	sell.Timestamp = time.Now()
	sell.Price = stockPrice
	userStock.updateStockAmount(sell.StockSoldAmount * -1)
	return nil
}

func (sell *Sell) commit() error {
	user := getUser(sell.UserId)
	user.updateUserBalance(sell.ActualCashAmount)
	_, err := sell.insertSell()
	return err
}

func (sell *Sell) cancel() {
	userStock := getOrCreateUserStock(sell.UserId, sell.StockSymbol)
	userStock.updateStockAmount(sell.StockSoldAmount)
}

func (sell *Sell) updateSell() error {
	_, err := db.Exec("update Sell set IntendedCashAmount=?, Price=?, ActualCashAmount=?, StockSoldAmount = ? where Id=?", sell.IntendedCashAmount, sell.Price, sell.ActualCashAmount, sell.StockSoldAmount, sell.Id)
	if err != nil {
		return err
	}
	return err
}

func (sell *Sell) insertSell() (*Sell, error) {
	res, err := db.Exec("insert into Sell(Price,StockSymbol,UserId,IntendedCashAmount,ActualCashAmount,StockSoldAmount) values(?,?,?,?,?,?)", sell.Price, sell.StockSymbol, sell.UserId, sell.IntendedCashAmount, sell.ActualCashAmount, sell.StockSoldAmount)
	if err != nil {
		return sell, err
	}
	sell.Id, err = res.LastInsertId()
	return sell, err
}

func getSell(id int64) *Sell {
	sell := &Sell{}
	err := db.QueryRow("Select * from Sell where Id=?", id).Scan(&sell.Id, &sell.Price, &sell.StockSymbol, &sell.UserId, &sell.IntendedCashAmount, &sell.ActualCashAmount, &sell.StockSoldAmount)
	if err != nil {
		log.Println(err)
	}
	return sell
}
