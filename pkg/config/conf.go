package config

import (
	"neko-acm/pkg/utils"
)

type Config struct {
	Server ServerConf `yaml:"server" json:"server"`
	Grpc   GrpcConf   `yaml:"grpc" json:"grpc"`
	Openai OpenaiConf `yaml:"openai" json:"openai"`
}

// InitConfig 初始化
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
	c.Server.Default()
	c.Grpc.Default()
	c.Openai.Default()
}
