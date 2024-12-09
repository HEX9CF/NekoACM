package bootstrap

import (
	"log"
	"neko-acm-ai/server"
)

func initServer() error {
	err := server.InitServer()
	if err != nil {
		log.Println("Init server failed!")
		return err
	}

	log.Println("Init server success.")
	return nil
}
