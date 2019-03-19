package main

import (
	"database/sql"
	"time"
	"os"
)

var createSellTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS ` + DB_NAME,
	`USE ` + DB_NAME,
	`CREATE TABLE IF NOT EXISTS Sell(
		Id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		Price float DEFAULT 0,
		StockSymbol varchar(3) NULL,
		UserId varchar(32) NULL,
		IntendedCashAmount float DEFAULT 0,
		ActualCashAmount float DEFAULT 0,
		StockSoldAmount int DEFAULT 0,
		PRIMARY KEY (Id)
	)`,
	`CREATE TABLE IF NOT EXISTS Sell_Trigger(
		UserId varchar(32),
		SellId INT UNSIGNED,
		Active BOOLEAN DEFAULT false,
		UNIQUE(UserId, SellId)
	)`,
}

var createUserTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS ` + DB_NAME,
	`USE ` + DB_NAME,
	`CREATE TABLE IF NOT EXISTS User(
		Id varchar(32) NOT NULL,
		Balance float DEFAULT 0, 
		Name varchar(32) NULL,
		PRIMARY KEY (id)
	)`,
	`CREATE TABLE IF NOT EXISTS User_Stock(
		UserId varchar(32),
		StockSymbol varchar(3) NULL,
		Amount INT UNSIGNED DEFAULT 0,
		UNIQUE(UserId, StockSymbol)
	)`,
}

var createBuyTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS ` + DB_NAME,
	`USE ` + DB_NAME,
	`CREATE TABLE IF NOT EXISTS Buy(
		Id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		Price float DEFAULT 0,
		StockSymbol varchar(3) NULL,
		UserId varchar(32) NULL,
		IntendedCashAmount float DEFAULT 0,
		ActualCashAmount float DEFAULT 0,
		StockBoughtAmount int DEFAULT 0,
		PRIMARY KEY (Id)
	)`,
	`CREATE TABLE IF NOT EXISTS Buy_Trigger(
		UserId varchar(32),
		BuyId INT UNSIGNED,
		Active BOOLEAN DEFAULT false,
		UNIQUE(UserId, BuyId)
	)`,
}
var sellDb *sql.DB
var buyDb *sql.DB
var userDb *sql.DB

func createAndOpenSellDB(){
	var err error
	sellDb, err = sql.Open("mysql", "root:@tcp("+os.Getenv("DAYTRADER_SELL_DB_IP")+":3306)/")
	if err != nil {
		panic(err)
	}

	for _, stmt := range createSellTableStatements {
		_, err := sellDb.Exec(stmt)
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
	}
	sellDb, err = sql.Open("mysql", "root:@tcp("+os.Getenv("DAYTRADER_SELL_DB_IP")+":3306)/daytrader")
	if err != nil {
		panic(err)
	}
	sellDb.SetConnMaxLifetime(time.Second * 0)
	sellDb.SetMaxIdleConns(10000)
}

func createAndOpenBuyDB(){
	var err error
	buyDb, err = sql.Open("mysql", "root:@tcp("+os.Getenv("DAYTRADER_BUY_DB_IP")+":3306)/")
	if err != nil {
		panic(err)
	}

	for _, stmt := range createBuyTableStatements {
		_, err := buyDb.Exec(stmt)
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
	}
	buyDb, err = sql.Open("mysql", "root:@tcp("+os.Getenv("DAYTRADER_BUY_DB_IP")+":3306)/daytrader")
	if err != nil {
		panic(err)
	}
	buyDb.SetConnMaxLifetime(time.Second * 0)
	buyDb.SetMaxIdleConns(10000)
}

func createAndOpenUserDB(){
	var err error
	userDb, err = sql.Open("mysql", "root:@tcp("+os.Getenv("DAYTRADER_USER_DB_IP")+":3306)/")
	if err != nil {
		panic(err)
	}

	for _, stmt := range createUserTableStatements {
		_, err := userDb.Exec(stmt)
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
	}
	userDb, err = sql.Open("mysql", "root:@tcp("+os.Getenv("DAYTRADER_USER_DB_IP")+":3306)/daytrader")
	if err != nil {
		panic(err)
	}
	userDb.SetConnMaxLifetime(time.Second * 0)
	userDb.SetMaxIdleConns(10000)
}

func createAndOpenDB() {
	createAndOpenBuyDB()
	createAndOpenSellDB()
	createAndOpenUserDB()
}
