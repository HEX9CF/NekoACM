package config

// NacosConf Nacos的整体配置
type NacosConf struct {
	Client NacosClientConf   `yaml:"client" json:"client"`
	Server []NacosServerConf `yaml:"server" json:"server"`
}

// NacosClientConf Nacos客户端配置
type NacosClientConf struct {
	NamespaceId         string `yaml:"namespace-id" json:"namespace_id"`
	TimeoutMs           uint64 `yaml:"timeout-ms" json:"timeout_ms"`
	NotLoadCacheAtStart bool   `yaml:"not-load-cache-at-start" json:"not_load_cache_at_start"`
	LogDir              string `yaml:"log-dir" json:"log_dir"`
	CacheDir            string `yaml:"cache-dir" json:"cache_dir"`
	LogLevel            string `yaml:"log-level" json:"log_level"`
}

// NacosServerConf Nacos服务器配置
type NacosServerConf struct {
	IpAddr      string `yaml:"ip-addr" json:"ip_addr"`
	Port        uint64 `yaml:"port" json:"port"`
	Scheme      string `yaml:"scheme" json:"scheme"`
	ContextPath string `yaml:"context-path" json:"context_path"`
}

// Default 为NacosConf设置默认值
func (n *NacosConf) Default() {
	n.Client.Default()
	n.Server[0].Default()
}

// Default 为NacosClientConf设置默认值
func (c *NacosClientConf) Default() {
	c.NamespaceId = "e525eafa-f7d7-4029-83d9-008937f9d468"
	c.TimeoutMs = 5000
	c.NotLoadCacheAtStart = true
	c.LogDir = "/tmp/nacos/log"
	c.CacheDir = "/tmp/nacos/cache"
	c.LogLevel = "debug"
}

// Default 为NacosServerConf设置默认值
func (s *NacosServerConf) Default() {
	s.IpAddr = "localhost"
	s.Port = 80
	s.Scheme = "http"
	s.ContextPath = "/nacos"
}
