package nacos

func InitNacos() error {
	initConfig()

	err := initClient()
	if err != nil {
		return err
	}

	return nil
}
