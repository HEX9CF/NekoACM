package main

import (
	"neko-acm/cmd/bootstrap"
	"neko-acm/cmd/cli"
)

func main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
	cli.Execute()
}
