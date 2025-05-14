package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func GetConfig() error {
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "dataId",
		Group:  "group"})
	if err != nil {
		return err
	}
	if content == "" {
		return errors.New("Nacos 获取配置失败")
	}

	return nil
}
