package main

import (
	"ecommerce/api/meli"
	"ecommerce/api/redis"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	ch := redis.NewCache(os.Getenv("CACHE_ADDR"))

	sr := meli.NewService(meli.Config{
		AppID:    os.Getenv("MELI_APP_ID"),
		Secret:   os.Getenv("MELI_SECRET"),
		Redirect: os.Getenv("MELI_REDIRECT"),
		Cache:    ch,
	})

	s := NewServer(sr, sr, ch)

	s.Init()
}
