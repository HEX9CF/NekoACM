package conf

type DatabaseConf struct {
	Host    string `yaml:"host" json:"host"`
	Port    string `yaml:"port" json:"port"`
	Name    string `yaml:"name" json:"name"`
	User    string `yaml:"user" json:"user"`
	Pwd     string `yaml:"password" json:"password"`
	MaxConn int    `yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdle int    `yaml:"max_idle_conns" json:"max_idle_conns"`
}

func (d *DatabaseConf) Default() {
	d.Host = "stuoj-ai-db"
	d.Port = "3306"
	d.Name = "stuoj-ai-db"
	d.User = "stuoj"
	d.Pwd = "stuoj"
	d.MaxConn = 10
	d.MaxIdle = 5
}
