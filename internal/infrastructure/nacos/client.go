package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"nekoacm-server/pkg/config"
)

var (
	clientConfig  constant.ClientConfig
	serverConfigs []constant.ServerConfig
	namingClient  naming_client.INamingClient
	configClient  config_client.IConfigClient
)

func initNacosConfig() {
	cConf := config.Conf.Nacos.Client
	sConfs := config.Conf.Nacos.Server

	// 创建clientConfig
	clientConfig = *constant.NewClientConfig(
		constant.WithNamespaceId(cConf.NamespaceId),
		constant.WithTimeoutMs(cConf.TimeoutMs),
		constant.WithNotLoadCacheAtStart(cConf.NotLoadCacheAtStart),
		constant.WithLogDir(cConf.LogDir),
		constant.WithCacheDir(cConf.CacheDir),
		constant.WithLogLevel(cConf.LogLevel),
	)

	// 创建serverConfig
	serverConfigs = make([]constant.ServerConfig, 0, len(sConfs))
	for _, conf := range sConfs {
		serverConfigs = append(serverConfigs, *constant.NewServerConfig(
			conf.IpAddr,
			conf.Port,
			constant.WithScheme(conf.Scheme),
			constant.WithContextPath(conf.ContextPath),
		))
	}
}

func initClient() error {
	var err error

	// 创建服务发现客户端
	namingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return err
	}

	// 创建动态配置客户端
	configClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
