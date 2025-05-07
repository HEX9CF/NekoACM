package config

type OpenaiConf struct {
	Model   string `yaml:"model" json:"model"`
	BaseUrl string `yaml:"base_url" json:"base_url"`
	ApiKey  string `yaml:"api_key" json:"api_key"`
}

func (a *OpenaiConf) Default() {
	a.Model = "gpt-3.5-turbo"
	a.BaseUrl = "https://api.openai.com"
	a.ApiKey = "sk-1234567890"
}
