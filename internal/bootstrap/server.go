package bootstrap

import (
	"log"
	"neko-acm-ai/server"
)

func initServer() {
	err := server.InitServer()
	if err != nil {
		log.Println("Init server failed!")
		panic(err)
	}

	log.Println("Init server success.")
}
