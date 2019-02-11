package main

import (
	"database/sql"
	"time"

	pb "./protobuff"
)

type mockCache struct{}
type mockDB struct{}
type mockResult struct{}

var s server
var symbol = "ABC"
var userId = "tester"
var filename = "dumplog"
var amount = float32(521)
var transactionId = int32(1)
var quotePrice = float32(5.21)
var hash = "lod23EP0lofFCkEd0ilcUpjL0MuBcIh3HiwAq9QSXdU="
var name = "tester"
var userCache = make(map[string]*User)
var stockCache = make(map[string]*Stock)

func genGrpcRequest(name string) *pb.Command {
	req := &pb.Command{UserId: userId, Symbol: symbol, Amount: amount, TransactionId: transactionId, Name: name, Filename: filename}
	return req
}

func (c *mockCache) setCache(key string, value interface{}) error {
	if v, ok := value.(*User); ok {
		userCache[key] = v
	} else if v, ok := value.(*Stock); ok {
		stockCache[key] = v
	}
	return nil
}

func (c *mockCache) getCacheUser(key string) (*User, error) {
	return userCache[key], nil
}

func (c *mockCache) getCacheStock(key string) (*Stock, error) {
	return stockCache[key], nil
}

func (db *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	res := mockResult{}
	return res, nil
}

func (m mockResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (m mockResult) RowsAffected() (int64, error) {
	return 0, nil
}

func (db *mockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	row := &sql.Row{}
	return row
}

func (db *mockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows := &sql.Rows{}
	return rows, nil
}

func (db *mockDB) SetConnMaxLifetime(d time.Duration) {
	return
}

func (db *mockDB) SetMaxIdleConns(n int) {
	return
}
