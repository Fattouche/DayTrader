package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"
)

func (sell *Sell) toString() string {
	if sell == nil {
		return ""
	}
	bytes, _ := sell.MarshalJSON()
	return string(bytes)
}

func createSell(ctx context.Context, intendedCashAmount float32, symbol string, user *User, writeThrough bool) (*Sell, error) {
	stock, err := quote(ctx, user.Id, symbol)
	if err != nil {
		return nil, err
	}
	sell := &Sell{Price: stock.Price, StockSymbol: symbol, UserId: user.Id}
	err = sell.updateCashAmount(ctx, intendedCashAmount, user)
	if err != nil {
		return nil, err
	}
	err = sell.updatePrice(ctx, stock.Price, user, writeThrough)
	if err != nil {
		return nil, err
	}
	return sell, err
}

func (sell *Sell) updateCashAmount(ctx context.Context, amount float32, user *User) error {
	stock, _ := quote(ctx, sell.UserId, sell.StockSymbol)
	userStock := getOrCreateUserStock(ctx, sell.UserId, sell.StockSymbol, user)
	stockSoldAmount := int(math.Floor(float64(amount / stock.Price)))
	if stockSoldAmount > userStock.Amount {
		return fmt.Errorf("Not enough stock, have %d need %d", userStock.Amount, stockSoldAmount)
	}
	sell.IntendedCashAmount = amount
	return nil
}

func (sell *Sell) updatePrice(ctx context.Context, stockPrice float32, user *User, writeThrough bool) error {
	userStock := getOrCreateUserStock(ctx, sell.UserId, sell.StockSymbol, user)
	updateSoldAmount := int(math.Min(math.Floor(float64(sell.IntendedCashAmount/stockPrice)), float64(userStock.Amount+sell.StockSoldAmount)))
	updated := updateSoldAmount - sell.StockSoldAmount
	sell.StockSoldAmount += updated
	sell.ActualCashAmount = float32(sell.StockSoldAmount) * stockPrice
	sell.Timestamp = time.Now()
	sell.Price = stockPrice
	userStock.updateStockAmount(ctx, updated*-1, user, writeThrough)
	return nil
}

func (sell *Sell) commit(ctx context.Context, update bool, user *User) *User {
	err := user.updateStockBalance(ctx, sell.StockSymbol)
	if err != nil {
		return nil
	}
	user.updateUserBalance(ctx, sell.ActualCashAmount, true)
	sell.Committed = true
	if update {
		go sell.updateSell(ctx)
	} else {
		go sell.insertSell(ctx)
	}
	return user
}

func (sell *Sell) cancel(ctx context.Context, user *User, writeThrough bool) {
	userStock := getOrCreateUserStock(ctx, sell.UserId, sell.StockSymbol, user)
	userStock.updateStockAmount(ctx, sell.StockSoldAmount, user, writeThrough)
}

func (sell *Sell) updateSell(ctx context.Context) error {
	_, err := db.Exec("update Sell set IntendedCashAmount=?, Price=?, ActualCashAmount=?, StockSoldAmount = ?, Committed=? where Id=?", sell.IntendedCashAmount, sell.Price, sell.ActualCashAmount, sell.StockSoldAmount, sell.Committed, sell.Id)
	if err != nil {
		return err
	}
	return err
}

func (sell *Sell) insertSell(ctx context.Context) (*Sell, error) {
	res, err := db.Exec("insert into Sell(Price,StockSymbol,UserId,IntendedCashAmount,ActualCashAmount,StockSoldAmount, Committed) values(?,?,?,?,?,?, true)", sell.Price, sell.StockSymbol, sell.UserId, sell.IntendedCashAmount, sell.ActualCashAmount, sell.StockSoldAmount)
	if err != nil {
		return sell, err
	}
	sell.Id, err = res.LastInsertId()
	return sell, err
}

func getSell(ctx context.Context, id int64) *Sell {
	sell := &Sell{}
	err := db.QueryRow("Select * from Sell where Id=?", id).Scan(&sell.Id, &sell.Price, &sell.StockSymbol, &sell.UserId, &sell.IntendedCashAmount, &sell.ActualCashAmount, &sell.StockSoldAmount, &sell.FromTrigger, &sell.Committed)
	if err != nil {
		log.Println(err)
	}
	return sell
}

func (sell *Sell) isExpired() bool {
	duration := time.Since(sell.Timestamp)
	if duration > time.Second*60 {
		return true
	}
	return false
}

func upsertSellTrigger(ctx context.Context, req *pb.Command, user *User) (*Sell, error) {
	sell := &Sell{StockSymbol: req.Symbol, UserId: req.UserId}
	err := sell.updateCashAmount(ctx, req.Amount, user)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("insert into Sell(Price,StockSymbol,UserId,IntendedCashAmount,ActualCashAmount,StockSoldAmount,FromTrigger,Committed) values(?,?,?,?,?,?,true,false) on duplicate key update IntendedCashAmount=?, ActualCashAmount=?,StockSoldAmount=?", sell.Price, sell.StockSymbol, sell.UserId, sell.IntendedCashAmount, sell.ActualCashAmount, sell.StockSoldAmount, sell.IntendedCashAmount, sell.ActualCashAmount, sell.StockSoldAmount)
	if err != nil {
		return nil, err
	}
	return sell, nil
}

func getSellTrigger(ctx context.Context, symbol, userId string) (*Sell, error) {
	sell := &Sell{}
	err := db.QueryRow("Select * from Sell where UserId=? and StockSymbol=? and FromTrigger=true and Committed=false", userId, symbol).Scan(&sell.Id, &sell.Price, &sell.StockSymbol, &sell.UserId, &sell.IntendedCashAmount, &sell.ActualCashAmount, &sell.StockSoldAmount, &sell.FromTrigger, &sell.Committed)
	if err != nil {
		return nil, err
	}
	return sell, nil
}

func setSellTriggerPrice(ctx context.Context, user *User, req *pb.Command) (*Sell, error) {
	sell, err := getSellTrigger(ctx, req.Symbol, req.UserId)
	if err != nil {
		return nil, err
	}
	sell.updatePrice(ctx, req.Amount, user, true)
	sell.updateSell(ctx)
	return sell, nil
}

func cancelSellTrigger(ctx context.Context, req *pb.Command, user *User) error {
	sell, err := getSellTrigger(ctx, req.Symbol, req.UserId)
	if err != nil {
		return err
	}
	sell.cancel(ctx, user, true)
	db.Exec("DELETE From Sell where UserId=? and StockSymbol=? and FromTrigger=true and Committed=false", req.UserId, req.Symbol)
	return nil
}

func checkSellTriggers() {
	rows, err := db.Query("SELECT * from Sell where Committed=false and FromTrigger=true")
	if err != nil {
		log.Println(err)
	}
	sells := make([]*Sell, 0)
	for rows.Next() {
		sell := &Sell{}
		err = rows.Scan(&sell.Id, &sell.Price, &sell.StockSymbol, &sell.UserId, &sell.IntendedCashAmount, &sell.ActualCashAmount, &sell.StockSoldAmount, &sell.FromTrigger, &sell.Committed)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		sells = append(sells, sell)
	}
	rows.Close()
	for _, sell := range sells {
		stock, _ := quote(context.Background(), sell.UserId, sell.StockSymbol)
		if sell.Price <= stock.Price {
			user := getUser(sell.UserId)
			sell.Committed = true
			sell.updatePrice(context.Background(), stock.Price, user, true)
			sell.updateSell(context.Background())
			user.updateUserBalance(context.Background(), sell.ActualCashAmount, true)
		}
	}
}
