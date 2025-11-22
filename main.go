package main

import (
	"nekoacm/cmd/cli"
	"nekoacm/internal/bootstrap"
)

func main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
	cli.Execute()
}
