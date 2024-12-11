package main

import (
	"neko-acm/cmd"
	"neko-acm/internal/bootstrap"
)

func main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
