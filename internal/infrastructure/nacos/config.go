package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"nekoacm-server/pkg/config"
	"nekoacm-server/pkg/utils"
)

func LoadConfig() error {
	var err error
	conf := config.Conf.Nacos.Config

	err = getConfig(&config.Conf.Grpc, conf.GrpcDataId)
	if err != nil {
		return err
	}

	err = getConfig(&config.Conf.Openai, conf.OpenaiDataId)
	if err != nil {
		return err
	}

	return nil
}

func getConfig(v interface{}, dataId string) error {
	conf := config.Conf.Nacos.Config

	log.Println("获取配置:", conf.Group, dataId)
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  conf.Group})
	if err != nil {
		return err
	}
	if content == "" {
		return nil
	}

	err = utils.UnmarshalYaml(v, content)
	if err != nil {
		return err
	}

	return nil
}
