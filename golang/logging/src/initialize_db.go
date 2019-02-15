package main

import (
	"database/sql"
	"time"
)

var (
	dbName = "logging" // Name of database to store logs
	db     *sql.DB
)

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS ` + dbName,
	`USE ` + dbName,
	`CREATE TABLE IF NOT EXISTS UserCommandLog(
		Timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		Server VARCHAR(32) NOT NULL,
		TransactionNum BIGINT NOT NULL,
		Command VARCHAR(32) NOT NULL,
		Username VARCHAR(32),
		StockSymbol VARCHAR(3),
		Filename VARCHAR(255),
		Funds DECIMAL(32, 2)
	)`,
	`CREATE TABLE IF NOT EXISTS QuoteServerLog(
		Timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Server VARCHAR(32) NOT NULL,
		TransactionNum BIGINT NOT NULL,
		Price DECIMAL(32, 2) NOT NULL,
		StockSymbol VARCHAR(3) NOT NULL,
		Username VARCHAR(32) NOT NULL,
		QuoteServerTime BIGINT NOT NULL,
		CryptoKey VARCHAR(128) NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS AccountTransactionLog(
		Timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Server VARCHAR(32) NOT NULL,
		TransactionNum BIGINT NOT NULL,
		Action VARCHAR(32) NOT NULL,
		Username VARCHAR(32) NOT NULL,
		Funds DECIMAL(32, 2) NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS SystemEventLog(
		Timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Server VARCHAR(32) NOT NULL,
		TransactionNum BIGINT NOT NULL,
		Command VARCHAR(16) NOT NULL,
		Username VARCHAR(32),
		StockSymbol VARCHAR(3),
		Filename VARCHAR(255),
		Funds DECIMAL(32, 2)
	)`,
	`CREATE TABLE IF NOT EXISTS ErrorEventLog(
		Timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Server VARCHAR(32) NOT NULL,
		TransactionNum BIGINT NOT NULL,
		Command VARCHAR(16) NOT NULL,
		Username VARCHAR(32),
		StockSymbol VARCHAR(3),
		Filename VARCHAR(255),
		Funds DECIMAL(32, 2),
		ErrorMessage TINYTEXT
	)`,
	`CREATE TABLE IF NOT EXISTS DebugEventLog(
		Timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Server VARCHAR(32) NOT NULL,
		TransactionNum BIGINT NOT NULL,
		Command VARCHAR(16) NOT NULL,
		Username VARCHAR(32),
		StockSymbol VARCHAR(3),
		Filename VARCHAR(255),
		Funds DECIMAL(32, 2),
		DebugMessage TINYTEXT
	)`,
}

func createAndOpenDB() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(logging_db:3306)/")
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
	db, err = sql.Open("mysql", "root@tcp(logging_db:3306)/logging")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Second * 0)
	db.SetMaxIdleConns(10000)
}
