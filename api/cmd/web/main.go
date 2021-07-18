package main

import (
	"ecommerce/api/meli"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	sr := meli.NewService(os.Getenv("MELI_APP_ID"), os.Getenv("MELI_SECRET"))

	s := NewServer(sr)

	s.Init()
}
