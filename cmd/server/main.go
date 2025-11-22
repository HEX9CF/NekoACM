package main

import (
	"nekoacm/internal/bootstrap"
)

func main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
	bootstrap.InitServer()
}
