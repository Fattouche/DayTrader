package main

import (
	"encoding/json"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

var cache *memcache.Client

func initCache() {
	cache = memcache.New("cache:11211")
}

func setCache(key string, val interface{}) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		log.Println("Error converting value to byte array")
	}
	item := &memcache.Item{
		Key:   key,
		Value: bytes,
	}
	return cache.Set(item)
}

func getCacheStock(key string) (*Stock, error) {
	var stock Stock
	item, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(item.Value, &stock)
	return &stock, err
}

func getCacheUser(key string) (*User, error) {
	user := &User{Id: key}
	item, err := cache.Get(key)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(item.Value, user)
	return user, err
}
