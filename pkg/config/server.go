package config

type ServerConf struct {
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (s *ServerConf) Default() {
	s.Port = "14515"
	s.Token = "token"
}
