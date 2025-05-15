package nacos

import "nekoacm-server/pkg/config"

func LoadConfig() error {
	var err error

	if err = NacosClient.GetConfig(&config.Conf.Grpc, "nekoacm-server-grpc.yaml"); err != nil {
		return err
	}

	if err = NacosClient.GetConfig(&config.Conf.Openai, "nekoacm-server-openai.yaml"); err != nil {
		return err
	}

	return nil
}
