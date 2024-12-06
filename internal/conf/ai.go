package conf

type AiConf struct {
	Provider string `yaml:"provider" json:"provider"`
	Model    string `yaml:"model" json:"model"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Key      string `yaml:"key" json:"key"`
}

func (a *AiConf) Default() {
	a.Provider = "openai"
	a.Model = "gpt-3.5-turbo"
	a.Host = ""
	a.Port = ""
	a.Key = ""
}
