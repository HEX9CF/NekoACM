package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"nekoacm-server/pkg/config"
)

var (
	clientConfig  constant.ClientConfig
	serverConfigs []constant.ServerConfig
)

func initConfig() {
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
