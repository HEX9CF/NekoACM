package nacos

import (
	"nekoacm-common/infrastructure/nacos"
	"nekoacm-server/pkg/config"
)

var (
	NacosClient *nacos.NacosClient
)

// InitNacos 初始化Nacos服务
func InitNacos() error {
	conf := &config.Conf.Nacos

	// 使用server包中的配置创建nacos客户端
	NacosClient = nacos.NewNacosClient(conf)

	// 初始化nacos客户端
	if err := NacosClient.Init(); err != nil {
		return err
	}

	return nil
}
