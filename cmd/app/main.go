package app

import "nekoacm-server/cmd/bootstrap"

func Main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
}
