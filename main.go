package main

import (
	"neko-acm/cmd"
	"neko-acm/internal/bootstrap"
)

func main() {
	bootstrap.Init()
	cmd.Execute()
}
