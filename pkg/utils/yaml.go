package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

// 读取yaml文件
func ReadYaml(v interface{}, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

// 写入yaml文件
func WriteYaml(v interface{}, path string) error {
	yamlData, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

// 解码yaml字符串
func UnmarshalYaml(v interface{}, content string) error {
	var err error

	err = yaml.Unmarshal([]byte(content), v)
	if err != nil {
		return err
	}

	return nil
}
