package nacos

func InitNacos() error {
	var err error

	initNacosConfig()

	err = initClient()
	if err != nil {
		return err
	}

	return nil
}
