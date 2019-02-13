package main

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

//easyjson:json
type Stock struct {
	Symbol    string
	Price     float32
	Hash      string
	TimeStamp time.Time
}

//easyjson:json
type User struct {
	Balance   float32
	Name      string
	Id        string
	BuyStack  []*Buy
	SellStack []*Sell
}

//easyjson:json
type Buy struct {
	Id                 int64
	Price              float32
	StockSymbol        string
	IntendedCashAmount float32
	ActualCashAmount   float32
	StockBoughtAmount  int
	UserId             string
	Timestamp          time.Time
}

//easyjson:json
type Sell struct {
	Id                 int64
	Price              float32
	StockSymbol        string
	IntendedCashAmount float32
	ActualCashAmount   float32
	StockSoldAmount    int
	UserId             string
	Timestamp          time.Time
}

var cache *memcache.Client

func initCache() {
	cache = memcache.New("cache:11211")
}

func setCache(key string, val interface{}) error {
	var bytes []byte
	if user, ok := val.(*User); ok {
		bytes, _ = user.MarshalJSON()
	}
	if stock, ok := val.(*Stock); ok {
		bytes, _ = stock.MarshalJSON()
	}
	item := &memcache.Item{
		Key:   key,
		Value: bytes,
	}
	return cache.Set(item)
}

func getCacheStock(key string) (*Stock, error) {
	stock := &Stock{}
	item, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	err = stock.UnmarshalJSON(item.Value)
	return stock, err
}

func getCacheUser(key string) (*User, error) {
	user := &User{}
	item, err := cache.Get(key)
	if err != nil {
		return user, err
	}
	err = user.UnmarshalJSON(item.Value)
	return user, err
}
