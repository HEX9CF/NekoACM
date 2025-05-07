package problem

import (
	"neko-acm/internal/model"
	"neko-acm/pkg/utils"
	"strconv"
	"time"
)

// 保存到json文件
func SaveJson(p model.Problem) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/problem/" + p.Title + "_" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(p, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
