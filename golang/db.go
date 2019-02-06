package main

import (
	"database/sql"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var db *sql.DB

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS ` + DB_NAME,
	`USE ` + DB_NAME,
	`CREATE TABLE IF NOT EXISTS User(
		id varchar(32) NOT NULL,
		balance float DEFAULT 0, 
		name varchar(32) NULL,
		PRIMARY KEY (id)
	)`,
	`CREATE TABLE IF NOT EXISTS Sell(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		price float DEFAULT 0,
		stock_symbol varchar(3) NULL,
		user_id varchar(32) NOT NULL,
		intended_cash_amount float DEFAULT 0,
		actual_cash_amount float DEFAULT 0,
		stock_sold_amount int DEFAULT 0,
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE
	)`,
	`CREATE TABLE IF NOT EXISTS Buy(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		price float DEFAULT 0,
		stock_symbol varchar(3) NULL,
		user_id varchar(32) NOT NULL,
		intended_cash_amount float DEFAULT 0,
		actual_cash_amount float DEFAULT 0,
		stock_bought_amount int DEFAULT 0,
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE
	)`,
	`CREATE TABLE IF NOT EXISTS Sell_Trigger(
		user_id varchar(32) NOT NULL,
		sell_id INT UNSIGNED NOT NULL,
		active BOOLEAN DEFAULT false,
		UNIQUE(user_id, sell_id),
		FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE,
		FOREIGN KEY (sell_id) REFERENCES Sell(id) ON DELETE CASCADE
	)`,
	`CREATE TABLE IF NOT EXISTS Buy_Trigger(
		user_id varchar(32) NOT NULL,
		buy_id INT UNSIGNED NOT NULL,
		active BOOLEAN DEFAULT false,
		UNIQUE(user_id, buy_id),
		FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE,
		FOREIGN KEY (buy_id) REFERENCES Buy(id) ON DELETE CASCADE
	)`,
}

type Stock struct {
	Symbol string
	Price  float32
	Hash   string
}

type Buy struct {
	id                   int64
	price                float32
	stock_symbol         string
	intended_cash_amount float32
	actual_cash_amount   float32
	stock_bought_amount  int
	user_id              string
}

type Sell struct {
	id                   int64
	price                float32
	stock_symbol         string
	intended_cash_amount float32
	actual_cash_amount   float32
	stock_sold_amount    int
	user_id              string
}

type SellTrigger struct {
	user_id string
	sell_id int64
	active  bool
}

type BuyTrigger struct {
	user_id string
	buy_id  int64
	active  bool
}

type User struct {
	Balance   float32
	Name      string
	Id        string
	BuyStack  []Buy
	SellStack []Sell
}

func quote(userID, symbol string) (*Stock, error) {
	stock, err := getCacheStock(symbol)
	if err != nil {
		//TODO: Do something with this hash
		stock = &Stock{Symbol: symbol}
		stock.Price, stock.Hash, err = executeRequest(userID, symbol)
		setCache(stock.Symbol, stock)
		return stock, err
	}
	return stock, err
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

func getUser(userID string) (*User, error) {
	user := &User{Id: userID}
	err := db.QueryRow("SELECT balance from User where id=?", user.Id).Scan(&user.Balance)
	if err != nil {
		return nil, err
	}
	temp, _ := getCacheUser(user.Id)
	if temp != nil {
		user.SellStack = temp.SellStack
		user.BuyStack = temp.BuyStack
	} else {
		user.SellStack = make([]Sell, 0)
		user.BuyStack = make([]Buy, 0)
		setCache(userID, user)
	}
	return user, nil
}

func createUser(userID string) error {
	_, err := db.Exec("insert ignore into User(id) values(?)", userID)
	return err
}

func createBuy(price, intendedCashAmount, actualCashAmount float32, stockBoughtAmount int, symbol, userID string) (*Buy, error) {
	buy := &Buy{price: price, intended_cash_amount: intendedCashAmount, actual_cash_amount: actualCashAmount, stock_bought_amount: stockBoughtAmount, stock_symbol: symbol, user_id: userID}
	res, err := db.Exec("insert into Buy(price,stock_symbol,user_id,intended_cash_amount,actual_cash_amount,stock_bought_amount) values(?,?,?,?,?,?)", price, symbol, userID, intendedCashAmount, actualCashAmount, stockBoughtAmount)
	if err != nil {
		return buy, err
	}
	buy.id, err = res.LastInsertId()
	return buy, err
}

func getBuys(userID, symbol string) ([]*Buy, error) {
	rows, err := db.Query("SELECT * from Buy where user_id=? and stock_symbol=?", userID, symbol)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	buys := make([]*Buy, len(cols))
	rows.Scan(buys)
	return buys, nil
}

func createSell(price, intendedCashAmount, actualCashAmount float32, stockSoldAmount int, symbol, userID string) (*Sell, error) {
	sell := &Sell{price: price, intended_cash_amount: intendedCashAmount, actual_cash_amount: actualCashAmount, stock_sold_amount: stockSoldAmount, stock_symbol: symbol, user_id: userID}
	res, err := db.Exec("insert into Buy(price,stock_symbol,user_id,intended_cash_amount,actual_cash_amount,stock_sold_amount) values(?,?,?,?,?,?)", price, symbol, userID, intendedCashAmount, actualCashAmount, stockSoldAmount)
	if err != nil {
		return sell, err
	}
	sell.id, err = res.LastInsertId()
	return sell, err
}

func getSells(userID, symbol string) ([]*Sell, error) {
	rows, err := db.Query("SELECT * from Sell where user_id=? and stock_symbol=?", userID, symbol)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	sells := make([]*Sell, len(cols))
	rows.Scan(sells)
	return sells, nil
}

func getBuyTrigger(userID, symbol string) (*BuyTrigger, error) {
	buyTrigger := &BuyTrigger{user_id: userID}
	err := db.QueryRow("SELECT * from Buy_Trigger where user_id=?", buyTrigger.user_id).Scan(&buyTrigger)
	return buyTrigger, err
}

func getSellTrigger(userID, symbol string) (*SellTrigger, error) {
	sellTrigger := &SellTrigger{user_id: userID}
	err := db.QueryRow("SELECT * from Sell_Trigger where user_id=?", sellTrigger.user_id).Scan(&sellTrigger)
	return sellTrigger, err
}

func createBuyTrigger(userID, symbol string, amount float32) (*BuyTrigger, error) {
	buy, err := createBuy(0, amount, 0, 0, symbol, userID)
	_, err = db.Exec("insert into Buy_Trigger(user_id,buy_id) values(?,?)", userID, buy.id)
	buyTrigger := &BuyTrigger{user_id: userID, buy_id: buy.id, active: false}
	return buyTrigger, err
}

func createSellTrigger(userID, symbol string, amount float32) (*SellTrigger, error) {
	sell, err := createSell(0, amount, 0, 0, symbol, userID)
	_, err = db.Exec("insert into Sell_Trigger(user_id,buy_id) values(?,?)", userID, sell.id)
	sellTrigger := &SellTrigger{user_id: userID, sell_id: sell.id, active: false}
	return sellTrigger, err
}

func (user User) updateUserBalance(amount float32) (User, error) {
	user.Balance += amount
	_, err := db.Exec("update User set balance=? where id=?", user.Balance, user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func createAndOpenDB() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}

	for _, stmt := range createTableStatements {
		_, err := db.Exec(stmt)
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
	}
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/daytrader")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Second * 0)
	db.SetMaxIdleConns(10000)
}
