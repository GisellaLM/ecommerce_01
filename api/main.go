package main

import "ecommerce/api/src"

func main() {
	sr := &src.InMemoryService{}
	sr.Init()

	s := src.NewServer(sr)

	s.Init()
}
