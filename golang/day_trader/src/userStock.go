package main

import (
	"context"
)

func (userStock *UserStock) toString() string {
	if userStock == nil {
		return ""
	}
	bytes, _ := userStock.MarshalJSON()
	return string(bytes)
}

func getAllUserStock(userID string) map[string]int32 {
	rows, err := db.Query("SELECT StockSymbol, Amount from User_Stock where UserId=?", userID)
	stocks := make(map[string]int32)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var stockSymbol string
			var amount int32
			if rows.Scan(&stockSymbol, &amount) == nil {
				stocks[stockSymbol] = amount
			}
		}
	}
	return stocks
}

func getOrCreateUserStock(ctx context.Context, userID, symbol string, user *User) *UserStock {
	if amount, ok := user.StockMap[symbol]; ok {
		return &UserStock{UserId: userID, StockSymbol: symbol, Amount: int(amount)}
	}
	amount := 0
	db.QueryRow("Select Amount from User_Stock where UserId=? and StockSymbol=?", user.Id, symbol).Scan(&amount)
	userStock := &UserStock{UserId: userID, StockSymbol: symbol, Amount: amount}
	user.setCache()
	return userStock
}

func (userStock *UserStock) updateStockAmount(ctx context.Context, amount int, user *User) {
	userStock.Amount += amount
	user.StockMap[userStock.StockSymbol] = int32(userStock.Amount)
	user.setCache()
}
