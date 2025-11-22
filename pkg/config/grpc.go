package config

type GrpcConf struct {
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (g *GrpcConf) Default() {
	g.Port = "14515"
	g.Token = "token"
}
