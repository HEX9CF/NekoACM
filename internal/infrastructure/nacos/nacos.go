package nacos

func InitNacos() error {
	var err error

	initConfig()

	err = initClient()
	if err != nil {
		return err
	}

	err = register()
	if err != nil {
		return err
	}

	return nil
}
