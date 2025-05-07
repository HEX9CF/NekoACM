package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// 写入json文件
func WriteJson(v interface{}, path string) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// 创建目录
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	// 创建文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入数据
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
