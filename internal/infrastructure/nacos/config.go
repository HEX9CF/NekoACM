package nacos

import "github.com/nacos-group/nacos-sdk-go/v2/common/constant"

var (
	clientConfig  constant.ClientConfig
	serverConfigs []constant.ServerConfig
)

func initConfig() {
	// 创建clientConfig
	clientConfig = *constant.NewClientConfig(
		constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建serverConfig
	serverConfigs = []constant.ServerConfig{
		*constant.NewServerConfig(
			"console1.nacos.io",
			80,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
}
