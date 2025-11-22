package app

import "nekoacm/cmd/bootstrap"

func Main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
}
