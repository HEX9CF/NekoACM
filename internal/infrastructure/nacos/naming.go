package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"nekoacm-server/pkg/config"
)

func Register() error {
	conf := config.Conf.Nacos.Register

	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          conf.Ip,
		Port:        conf.Port,
		ServiceName: conf.ServiceName,
		Weight:      conf.Weight,
		Enable:      conf.Enable,
		Healthy:     conf.Healthy,
	})
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Nacos 注册实例失败")
	}

	return err
}
