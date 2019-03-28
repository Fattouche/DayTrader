package main

import (
	"context"
	"log"
)

func (user *User) toString() string {
	if user == nil {
		return ""
	}
	bytes, _ := user.MarshalJSON()
	return string(bytes)
}

func (user *User) popFromBuyStack() *Buy {
	if len(user.BuyStack) == 0 {
		return nil
	}
	buy := user.BuyStack[len(user.BuyStack)-1]
	if buy.isExpired() {
		user.BuyStack = make([]*Buy, 0)
		buy = nil
	} else {
		user.BuyStack = user.BuyStack[:len(user.BuyStack)-1]
	}
	user.setCache()
	return buy
}

func (user *User) popFromSellStack() *Sell {
	if len(user.SellStack) == 0 {
		return nil
	}
	sell := user.SellStack[len(user.SellStack)-1]
	if sell.isExpired() {
		user.SellStack = make([]*Sell, 0)
		sell = nil
	} else {
		user.SellStack = user.SellStack[:len(user.SellStack)-1]
	}
	user.setCache()
	return sell
}

func userExists(userID string, password string) (float32, map[string]int32, error) {
	stocks := make(map[string]int32)
	var balance float32
	err := db.QueryRow("SELECT Balance from User where Id=? and Password=?", userID, password).Scan(&balance)
	if err == nil {
		getAllUserStock(userID, &stocks)
	}
	return balance, stocks, err
}

func getUser(userID string) *User {
	user, err := getCacheUser(userID)
	if err != nil {
		db.QueryRow("SELECT Balance from User where Id=?", user.Id).Scan(&user.Balance)
		user.SellStack = make([]*Sell, 0)
		user.BuyStack = make([]*Buy, 0)
		user.StockMap = make(map[string]int32)
		user.setCache()
	}
	return user
}

func createUser(userID string, password string) error {
	user := &User{Id: userID, Balance: 0, Password: password}
	user.SellStack = make([]*Sell, 0)
	user.BuyStack = make([]*Buy, 0)
	user.StockMap = make(map[string]int32)
	user.setCache()
	_, err := db.Exec("insert into User(Id,Password) values(?,?)", userID, password)
	return err
}

func (user *User) updateUserBalance(ctx context.Context, amount float32, writeThrough bool) (*User, error) {
	var accountAction string
	if amount < 0 {
		accountAction = "Remove"
	} else {
		accountAction = "Add"
	}
	user.Balance += amount
	//deal with max balances and overflow
	if user.Balance > MAX_BALANCE {
		user.Balance = MAX_BALANCE
	}
	//Shouldn't happen
	if user.Balance < 0 {
		user.Balance = 0
	}
	if writeThrough {
		_, err := db.Exec("update User set Balance=? where Id=?", user.Balance, user.Id)
		if err != nil {
			return user, err
		}
	}
	pbLog, err := makeLogFromContext(ctx)
	if err != nil {
		log.Println(err)
	}
	pbLog.Username = user.Id
	pbLog.AccountAction = accountAction
	pbLog.Funds = user.Balance
	logEvent := &logObj{log: &pbLog, funcName: "LogAccountTransaction"}
	logChan <- logEvent
	user.setCache()
	return user, nil
}

func (user *User) updateStockBalance(ctx context.Context, symbol string) error {
	amount := user.StockMap[symbol]
	_, err := db.Exec("insert into User_Stock(UserId,StockSymbol,Amount) Values(?,?,?) on duplicate key update Amount=?", user.Id, symbol, amount, amount)
	if err != nil {
		return err
	}
	return nil
}
