package config

// NacosConf Nacos的整体配置
type NacosConf struct {
	Enable   bool              `yaml:"enable" json:"enable"`
	Client   NacosClientConf   `yaml:"client" json:"client"`
	Server   []NacosServerConf `yaml:"server" json:"server"`
	Register NacosRegisterConf `yaml:"register" json:"register"`
	Config   NacosConfigConf   `yaml:"config" json:"config"`
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

// NacosRegisterConf Nacos服务注册配置
type NacosRegisterConf struct {
	Ip          string  `yaml:"ip" json:"ip"`
	Port        uint64  `yaml:"port" json:"port"`
	ServiceName string  `yaml:"service-name" json:"service_name"`
	Weight      float64 `yaml:"weight" json:"weight"`
	Enable      bool    `yaml:"enable" json:"enable"`
	Healthy     bool    `yaml:"healthy" json:"healthy"`
}

// NacosConfigConf Nacos配置中心配置
type NacosConfigConf struct {
	Group        string `yaml:"group" json:"group"`
	GrpcDataId   string `yaml:"grpc-data-id" json:"grpc_data_id"`
	OpenaiDataId string `yaml:"openai-data-id" json:"openai_data_id"`
}

// Default 为NacosConf设置默认值
func (n *NacosConf) Default() {
	n.Enable = false

	n.Client.Default()
	n.Server[0].Default()
	n.Register.Default()
	n.Config.Default()
}

// Default 为NacosClientConf设置默认值
func (c *NacosClientConf) Default() {
	c.NamespaceId = ""
	c.TimeoutMs = 5000
	c.NotLoadCacheAtStart = true
	c.LogDir = "/tmp/nacos/log"
	c.CacheDir = "/tmp/nacos/cache"
	c.LogLevel = "debug"
}

// Default 为NacosServerConf设置默认值
func (s *NacosServerConf) Default() {
	s.IpAddr = "localhost"
	s.Port = 8848
	s.Scheme = "http"
	s.ContextPath = "/nacos"
}

// Default 为NacosRegisterConf设置默认值
func (r *NacosRegisterConf) Default() {
	r.Ip = "127.0.0.1"
	r.Port = 14516
	r.ServiceName = "nekoacm-server"
	r.Weight = 10
	r.Enable = true
	r.Healthy = true
}

func (c *NacosConfigConf) Default() {
	c.Group = "DEFAULT_GROUP"
	c.GrpcDataId = "nekoacm-server-grpc.yaml"
	c.OpenaiDataId = "nekoacm-server-openai.yaml"
}
