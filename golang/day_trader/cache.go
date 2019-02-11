package main

import (
	"encoding/json"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

type cache struct {
	Client *memcache.Client
}

func (c *cache) setCache(key string, val interface{}) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		log.Println("Error converting value to byte array")
	}
	item := &memcache.Item{
		Key:   key,
		Value: bytes,
	}
	return c.Client.Set(item)
}

func (c *cache) getCacheStock(key string) (*Stock, error) {
	var stock Stock
	item, err := c.Client.Get(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(item.Value, &stock)
	return &stock, err
}

func (c *cache) getCacheUser(key string) (*User, error) {
	user := &User{Id: key}
	item, err := c.Client.Get(key)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(item.Value, user)
	return user, err
}
