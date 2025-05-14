package config

import (
	"gopkg.in/yaml.v3"
	"nekoacm-server/pkg/utils"
)

var (
	Conf = &Config{}
)

type Config struct {
	Grpc   GrpcConf   `yaml:"grpc" json:"grpc"`
	Openai OpenaiConf `yaml:"openai" json:"openai"`
	Nacos  NacosConf  `yaml:"nacos" json:"nacos"`
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

// LoadConfig 加载配置
func LoadConfig(content string) error {
	tempConf := Config{}

	err := yaml.Unmarshal([]byte(content), &tempConf)
	if err != nil {
		return err
	}

	Conf.Grpc = tempConf.Grpc
	Conf.Openai = tempConf.Openai

	return nil
}

func (c *Config) Default() {
	c.Grpc.Default()
	c.Openai.Default()
	c.Nacos.Default()
}
