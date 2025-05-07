package problem

import (
	"encoding/json"
	"errors"
	"log"
	"neko-acm/internal/infrastructure/open_ai"
	"neko-acm/internal/model"
	"neko-acm/prompt"
	"neko-acm/utils"
)

// 翻译题目
func Translate(pi model.TranslateInstruction) (model.Problem, error) {
	var p model.Problem

	// 题目说明转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return model.Problem{}, err
	}
	log.Println("请求翻译题目：" + instruction)

	// 请求模型
	resp, err := open_ai.Chat(prompt.ProblemTranslate, instruction)
	if err != nil {
		return model.Problem{}, errors.New("请求模型失败！")
	}
	log.Println("翻译结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		log.Println(err)
		return model.Problem{}, errors.New("解析结果失败，请重试！")
	}

	return p, nil
}
