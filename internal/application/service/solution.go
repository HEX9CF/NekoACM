package service

import (
	"encoding/json"
	"errors"
	"log"
	"nekoacm-common/pkg/utils"
	"nekoacm-server/internal/application/dto"
	"nekoacm-server/internal/infrastructure/openai"
	"nekoacm-server/prompt"
	"strconv"
	"time"
)

// 生成题解
func SolutionGenerate(si dto.SolutionInstruction) (dto.Solution, error) {
	var s dto.Solution

	// 说明转换为字符串
	instruction, err := utils.PrettyStruct(si)
	if err != nil {
		return dto.Solution{}, err
	}
	log.Println("请求生成题解：" + instruction)

	// 请求模型
	resp, err := openai.Chat(prompt.SolutionGenerate, instruction)
	if err != nil {
		log.Println(err)
		return dto.Solution{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		log.Println(err)
		return dto.Solution{}, errors.New("解析结果失败，请重试！")
	}

	return s, nil
}

// 保存到json文件
func SolutionSaveJson(s dto.Solution) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/solution/" + s.Language + "_" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(s, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
