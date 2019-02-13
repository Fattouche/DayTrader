package main

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var cache *memcache.Client

func initCache() {
	cache = memcache.New("cache:11211")
}

func (user *User) setCache() error {
	bytes, _ := user.MarshalJSON()
	item := &memcache.Item{
		Key:   user.Id,
		Value: bytes,
	}
	return cache.Set(item)
}

func (stock *Stock) setCache() error {
	bytes, _ := stock.MarshalJSON()
	item := &memcache.Item{
		Key:   stock.Symbol,
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
	user := &User{Id: key}
	item, err := cache.Get(key)
	if err != nil {
		return user, err
	}
	err = user.UnmarshalJSON(item.Value)
	return user, err
}
