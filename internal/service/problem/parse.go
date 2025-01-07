package problem

import (
	"encoding/json"
	"errors"
	"log"
	"neko-acm/external/open_ai"
	"neko-acm/internal/model"
	"neko-acm/prompt"
	"neko-acm/utils"
)

// 解析题目
func Parse(pi model.ProblemData) (model.Problem, error) {
	var p model.Problem

	// 题目数据转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return model.Problem{}, err
	}
	log.Println("请求解析题目：" + instruction)

	// 请求模型
	resp, err := open_ai.Chat(prompt.ParseSystem.String(), instruction)
	if err != nil {
		return model.Problem{}, errors.New("请求模型失败！")
	}
	log.Println("解析结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		log.Println(err)
		return model.Problem{}, errors.New("解析结果失败，请重试！")
	}

	return p, nil
}
