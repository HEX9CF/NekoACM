package config

import (
	"nekoacm-common/infrastructure/nacos"
	"nekoacm-common/pkg/config"
	"nekoacm-common/pkg/utils"
)

var (
	Conf = &Config{}
)

type Config struct {
	config.CommonConfig
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
	c.CommonConfig.Default()
	c.Grpc.Default()
	c.Openai.Default()
}

func LoadConfig() error {
	var err error
	conf := Conf.Nacos.Config

	err = nacos.GetConfig(&Conf.Grpc, conf.GrpcDataId)
	if err != nil {
		return err
	}

	err = nacos.GetConfig(&Conf.Openai, conf.OpenaiDataId)
	if err != nil {
		return err
	}

	return nil
}
