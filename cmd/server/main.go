package server

import (
	"nekoacm/internal/bootstrap"
)

func Main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
}
