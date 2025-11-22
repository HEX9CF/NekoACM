package main

import (
	"nekoacm/internal/bootstrap"
	"nekoacm/internal/interface/stdio"
)

func main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
	stdio.Execute()
}
