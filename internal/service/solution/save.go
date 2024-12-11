package solution

import (
	"neko-acm/internal/model"
	"neko-acm/utils"
	"strconv"
	"time"
)

func SaveJson(s model.Solution) (string, error) {
	timestamp := time.Now().Unix()
	path := "output/solution/" + s.Language + "_" + strconv.FormatInt(timestamp, 10) + ".json"
	err := utils.WriteJson(s, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
