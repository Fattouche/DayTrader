package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"
)

func (buy *Buy) toString() string {
	if buy == nil {
		return ""
	}
	bytes, _ := buy.MarshalJSON()
	return string(bytes)
}

func createBuy(ctx context.Context, intendedCashAmount float32, symbol string, user *User, writeThrough bool) (*Buy, error) {
	stock, err := quote(ctx, user.Id, symbol)
	if err != nil {
		return nil, err
	}
	buy := &Buy{Price: stock.Price, StockSymbol: symbol, UserId: user.Id}
	err = buy.updateCashAmount(ctx, intendedCashAmount, user, writeThrough)
	if err != nil {
		return nil, err
	}
	buy.updatePrice(stock.Price)
	buy.Timestamp = time.Now()
	user.BuyStack = append(user.BuyStack, buy)
	user.setCache()
	return buy, err
}

func (buy *Buy) updateCashAmount(ctx context.Context, amount float32, user *User, writeThrough bool) error {
	if amount > user.Balance {
		msg := fmt.Sprintf("Not enough balance, have %f need %f", user.Balance, amount)
		return errors.New(msg)
	}
	updatedAmount := buy.IntendedCashAmount - amount
	user.updateUserBalance(ctx, updatedAmount, writeThrough)
	buy.IntendedCashAmount = float32(math.Abs(float64(updatedAmount)))
	return nil
}

func (buy *Buy) updatePrice(stockPrice float32) {
	buy.Price = stockPrice
	buy.StockBoughtAmount = int(math.Floor(float64(buy.IntendedCashAmount / buy.Price)))
	buy.ActualCashAmount = float32(buy.StockBoughtAmount) * buy.Price
}

func (buy *Buy) commit(ctx context.Context, user *User, update bool) *UserStock {
	if update {
		go buy.updateBuy(ctx, true)
	} else {
		//log here instead maybe?
		go buy.insertBuy(ctx)
	}
	user.updateUserBalance(ctx, buy.IntendedCashAmount-buy.ActualCashAmount, true)
	userStock := getOrCreateUserStock(ctx, buy.UserId, buy.StockSymbol, user)
	userStock.updateStockAmount(ctx, buy.StockBoughtAmount, user, true)
	return userStock
}

func (buy *Buy) cancel(ctx context.Context, user *User, writeThrough bool) {
	user.updateUserBalance(ctx, buy.IntendedCashAmount, writeThrough)
}

func (buy *Buy) updateBuy(ctx context.Context, committed bool) error {
	_, err := db.Exec("update Buy set IntendedCashAmount=?, Price=?, ActualCashAmount=?, StockBoughtAmount = ?, Committed=? where Id=?", buy.IntendedCashAmount, buy.Price, buy.ActualCashAmount, buy.StockBoughtAmount, committed, buy.Id)
	if err != nil {
		return err
	}
	return err
}

func (buy *Buy) insertBuy(ctx context.Context) (*Buy, error) {
	res, err := db.Exec("insert into Buy(Price,StockSymbol,UserId,IntendedCashAmount,ActualCashAmount,StockBoughtAmount, Committed) values(?,?,?,?,?,?,true)", buy.Price, buy.StockSymbol, buy.UserId, buy.IntendedCashAmount, buy.ActualCashAmount, buy.StockBoughtAmount)
	if err != nil {
		return buy, err
	}
	buy.Id, err = res.LastInsertId()
	return buy, err
}

func getBuy(ctx context.Context, id int64) *Buy {
	buy := &Buy{}
	err := db.QueryRow("Select * from Buy where Id=?", id).Scan(&buy.Id, &buy.Price, &buy.StockSymbol, &buy.UserId, &buy.IntendedCashAmount, &buy.ActualCashAmount, &buy.StockBoughtAmount)
	if err != nil {
		log.Println(err)
	}
	return buy
}

func (buy *Buy) isExpired() bool {
	duration := time.Since(buy.Timestamp)
	if duration > time.Second*60 {
		return true
	}
	return false
}

func upsertBuyTrigger(ctx context.Context, req *pb.Command, user *User) (*Buy, error) {
	buy, err := createBuy(ctx, req.Amount, req.Symbol, user, true)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("insert into Buy(Price,StockSymbol,UserId,IntendedCashAmount,ActualCashAmount,StockBoughtAmount,FromTrigger,Committed) values(?,?,?,?,?,?,true,false) on duplicate key update IntendedCashAmount=?, ActualCashAmount=?,StockBoughtAmount=?", buy.Price, buy.StockSymbol, buy.UserId, buy.IntendedCashAmount, buy.ActualCashAmount, buy.StockBoughtAmount, buy.IntendedCashAmount, buy.ActualCashAmount, buy.StockBoughtAmount)
	if err != nil {
		return nil, err
	}
	return buy, nil
}

func getBuyTrigger(ctx context.Context, symbol, userId string) (*Buy, error) {
	buy := &Buy{}
	err := db.QueryRow("Select * from Buy where UserId=? and StockSymbol=? and FromTrigger=true and Committed=false", userId, symbol).Scan(&buy.Id, &buy.Price, &buy.StockSymbol, &buy.UserId, &buy.IntendedCashAmount, &buy.ActualCashAmount, &buy.StockBoughtAmount)
	if err != nil {
		return nil, err
	}
	return buy, nil
}

func setBuyTriggerPrice(ctx context.Context, req *pb.Command) (*Buy, error) {
	buy, err := getBuyTrigger(ctx, req.Symbol, req.UserId)
	if err != nil {
		return nil, err
	}
	buy.updatePrice(req.Amount)
	buy.updateBuy(ctx, false)
	return buy, nil
}

func cancelBuyTrigger(ctx context.Context, req *pb.Command, user *User) error {
	buy, err := getBuyTrigger(ctx, req.Symbol, req.UserId)
	if err != nil {
		return err
	}
	buy.cancel(ctx, user, true)
	db.Exec("DELETE From Buy where UserId=? and StockSymbol=? and FromTrigger=true and Committed=false", req.UserId, req.Symbol)
	return nil
}

func checkBuyTriggers() {
	rows, err := db.Query("SELECT * from Buy where FromTrigger=true and Committed=false")
	if err != nil {
		log.Println(err)
	}
	buys := make([]*Buy, 0)
	for rows.Next() {
		buy := &Buy{}
		err = rows.Scan(&buy.Id, &buy.StockSymbol, &buy.IntendedCashAmount, &buy.StockBoughtAmount, &buy.Price, &buy.UserId)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		buys = append(buys, buy)
	}
	rows.Close()
	for _, buy := range buys {
		stock, _ := quote(context.Background(), buy.UserId, buy.StockSymbol)
		if buy.Price >= stock.Price {
			buy.updatePrice(stock.Price)
			buy.commit(context.Background(), getUser(buy.UserId), true)
		}
	}
}
