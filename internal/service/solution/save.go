package solution

import (
	"neko-acm/internal/model"
	"neko-acm/pkg/utils"
	"strconv"
	"time"
)

// 保存到json文件
func SaveJson(s model.Solution) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/solution/" + s.Language + "_" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(s, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
