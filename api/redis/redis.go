package redis

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

const (
	cacheDuration = time.Hour * 24
)

type Cache struct {
	c *redis.Client
}

func NewCache(addr string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatalln("Cannot initialize redis cache")
		return nil
	}

	return &Cache{
		c: client,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	if c.c.Get(key).Err() == redis.Nil {
		c.c.Set(key, value, cacheDuration)
	}
}

func (c *Cache) Get(key string) (value interface{}) {
	err := c.c.Get(key).Err()

	if err == redis.Nil {
		return nil
	}

	c.c.Get(key).Scan(&value)

	return
}
