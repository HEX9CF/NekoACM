package main

import (
	"nekoacm-server/cmd/bootstrap"
	"nekoacm-server/cmd/cli"
)

func main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
	cli.Execute()
}
