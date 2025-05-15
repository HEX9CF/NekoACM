package nacos

import "nekoacm-server/pkg/config"

func LoadConfig() error {
	var err error
	conf := config.Conf.Nacos.Config

	if err = NacosClient.GetConfig(&config.Conf.Grpc, conf.GrpcDataId); err != nil {
		return err
	}

	if err = NacosClient.GetConfig(&config.Conf.Openai, conf.OpenaiDataId); err != nil {
		return err
	}

	return nil
}
