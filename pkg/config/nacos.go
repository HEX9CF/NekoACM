package config

// NacosConf Nacos的整体配置
type NacosConf struct {
	Client NacosClientConf   `yaml:"client" json:"client"`
	Server []NacosServerConf `yaml:"server" json:"server"`
}

// NacosClientConf Nacos客户端配置
type NacosClientConf struct {
	Namespace           string `yaml:"namespace" json:"namespace"`
	TimeoutMs           int    `yaml:"timeout-ms" json:"timeout-ms"`
	NotLoadCacheAtStart bool   `yaml:"not-load-cache-at-start" json:"not-load-cache-at-start"`
	LogDir              string `yaml:"log-dir" json:"log-dir"`
	CacheDir            string `yaml:"cache-dir" json:"cache-dir"`
	LogLevel            string `yaml:"log-level" json:"log-level"`
}

// NacosServerConf Nacos服务器配置
type NacosServerConf struct {
	IP          string `yaml:"ip" json:"ip"`
	Port        int    `yaml:"port" json:"port"`
	Scheme      string `yaml:"scheme" json:"scheme"`
	ContextPath string `yaml:"context-path" json:"context-path"`
}

// Default 为NacosConf设置默认值
func (n *NacosConf) Default() {
	n.Client.Default()
	n.Server[0].Default()
}

// Default 为NacosClientConf设置默认值
func (c *NacosClientConf) Default() {
	c.Namespace = "e525eafa-f7d7-4029-83d9-008937f9d468"
	c.TimeoutMs = 5000
	c.NotLoadCacheAtStart = true
	c.LogDir = "/tmp/nacos/log"
	c.CacheDir = "/tmp/nacos/cache"
	c.LogLevel = "debug"
}

// Default 为NacosServerConf设置默认值
func (s *NacosServerConf) Default() {
	s.IP = "localhost"
	s.Port = 80
	s.Scheme = "http"
	s.ContextPath = "/nacos"
}
