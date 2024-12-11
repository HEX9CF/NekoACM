package problem

import (
	"neko-acm/internal/model"
	"neko-acm/utils"
	"strconv"
	"time"
)

func SaveJson(p model.Problem) (string, error) {
	timestamp := time.Now().Unix()
	path := "output/problem/" + p.Title + "_" + strconv.FormatInt(timestamp, 10) + ".json"
	err := utils.WriteJson(p, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
