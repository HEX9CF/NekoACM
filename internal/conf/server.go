package conf

type ServerConf struct {
	Port string `yaml:"port" json:"port"`
}

func (s *ServerConf) Default() {
	s.Port = "14515"
}
