package bootstrap

// 初始化
func Init() error {
	if err := initConfig(); err != nil {
		return err
	}
	if err := initLlm(); err != nil {
		return err
	}
	return nil
}
