package conf

import (
	"stuoj-ai/utils"
)

type Config struct {
	Datebase DatabaseConf `yaml:"database" json:"database"`
	Server   ServerConf   `yaml:"server" json:"server"`
	Ai       AiConf       `yaml:"ai" json:"ai"`
}

// Config 初始化
func InitConfig() error {
	v, err := utils.IsFileExists("config.yaml")
	if err != nil {
		return err
	}
	if !v {
		Conf.Default()
		err = utils.WriteYaml(&Conf, "config.yaml")
		if err != nil {
			return err
		}
	}
	err = utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Default() {
	c.Datebase.Default()
	c.Server.Default()
	c.Ai.Default()
}
