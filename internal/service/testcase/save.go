package testcase

import (
	"neko-acm/internal/model"
	"neko-acm/utils"
	"strconv"
	"time"
)

func SaveJson(t model.Testcase) (string, error) {
	timestamp := time.Now().Unix()
	path := "output/testcase/" + strconv.FormatInt(timestamp, 10) + ".json"
	err := utils.WriteJson(t, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
