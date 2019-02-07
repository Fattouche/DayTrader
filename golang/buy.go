package main

import (
	"errors"
	"fmt"
	"math"
)

func createBuy(price, intendedCashAmount, actualCashAmount float32, stockBoughtAmount int, symbol, userID string) (*Buy, error) {
	buy := &Buy{price: price, intended_cash_amount: intendedCashAmount, actual_cash_amount: actualCashAmount, stock_bought_amount: stockBoughtAmount, stock_symbol: symbol, user_id: userID}
	err := buy.updateCashAmount(intendedCashAmount)
	if err != nil {
		return nil, err
	}
	buy.updatePrice(price)
	return buy, err
}

func (buy Buy) updateCashAmount(amount float32) error {
	user, _ := getUser(buy.user_id)
	if amount > user.Balance {
		msg := fmt.Sprintf("Not enough balance, have %f need %f", user.Balance, amount)
		return errors.New(msg)
	}
	updatedAmount := buy.intended_cash_amount - amount
	user.updateUserBalance(updatedAmount)
	buy.intended_cash_amount = float32(math.Abs(float64(updatedAmount)))
	return nil
}

func (buy Buy) updatePrice(stockPrice float32) {
	buy.price = stockPrice
	buy.stock_bought_amount = int(math.Floor(float64(buy.intended_cash_amount / buy.price)))
	buy.actual_cash_amount = float32(buy.stock_bought_amount) * buy.price
}

func (buy Buy) commit() (*UserStock, error) {
	_, err := buy.insertBuy()
	user_stock, err := getOrCreateUserStock(buy.user_id, buy.stock_symbol)
	user_stock.updateStockAmount(buy.stock_bought_amount)
	return user_stock, err
}

func (buy Buy) cancel() {
	user, _ := getUser(buy.user_id)
	user.updateUserBalance(buy.intended_cash_amount)
}
