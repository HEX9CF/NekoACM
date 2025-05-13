package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func InitNacos() error {
	var err error

	initConfig()

	err = initClient()
	if err != nil {
		return err
	}

	err = register()
	if err != nil {
		return err
	}

	return nil
}

func register() error {
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "nekoacm-server",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // 默认值DEFAULT
		GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Nacos 注册实例失败")
	}

	return err
}
