package bootstrap

func Init() error {
	if err := initConfig(); err != nil {
		return err
	}
	if err := initLlm(); err != nil {
		return err
	}
	//if err := initServer(); err != nil {
	//	return err
	//}
	return nil
}
