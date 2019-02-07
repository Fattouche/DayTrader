package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func createSell(price, intendedCashAmount, actualCashAmount float32, stockSoldAmount int, symbol, userID string) (*Sell, error) {
	sell := &Sell{price: price, intended_cash_amount: intendedCashAmount, actual_cash_amount: actualCashAmount, stock_sold_amount: stockSoldAmount, stock_symbol: symbol, user_id: userID}
	err := sell.updateCashAmount(intendedCashAmount)
	if err != nil {
		return nil, err
	}
	err = sell.updatePrice(price)
	if err != nil {
		return nil, err
	}
	return sell, err
}

func (sell *Sell) updateCashAmount(amount float32) error {
	stock, _ := quote(sell.user_id, sell.stock_symbol)
	userStock, _ := getOrCreateUserStock(sell.user_id, sell.stock_symbol)
	stockSoldAmount := int(math.Floor(float64(amount / stock.Price)))
	if stockSoldAmount > userStock.amount {
		return errors.New(fmt.Sprint("Not enough stock, have %f need %f", userStock.amount, stockSoldAmount))
	}
	sell.intended_cash_amount = amount
	return nil
}

func (sell *Sell) updatePrice(stockPrice float32) error {
	sell.cancel()
	userStock, _ := getOrCreateUserStock(sell.user_id, sell.stock_symbol)
	sell.stock_sold_amount = int(math.Min(math.Floor(float64(sell.intended_cash_amount/stockPrice)), float64(userStock.amount)))
	if sell.stock_sold_amount <= 0 {
		return errors.New("Update trigger price failed")
	}
	sell.actual_cash_amount = float32(sell.stock_sold_amount) * stockPrice
	sell.stock_sold_amount = int(float32(sell.stock_sold_amount) * stockPrice)
	sell.timestamp = time.Now()
	sell.price = stockPrice
	userStock.updateStockAmount(sell.stock_sold_amount * -1)
	return nil
}

func (sell *Sell) commit() error {
	user := getUser(sell.user_id)
	user.updateUserBalance(sell.actual_cash_amount)
	_, err := sell.insertSell()
	return err
}

func (sell *Sell) cancel() {
	userStock, _ := getOrCreateUserStock(sell.user_id, sell.stock_symbol)
	userStock.updateStockAmount(sell.stock_sold_amount)
}
