package main

import (
	"github.com/bradfitz/gomemcache/memcache"
)

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
