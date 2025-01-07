package problem

import (
	"encoding/json"
	"errors"
	"log"
	"neko-acm/external/open_ai"
	"neko-acm/internal/model"
	"neko-acm/prompt"
)

// 解析题目
func Parse(pd model.ProblemData) (model.Problem, error) {
	var p model.Problem

	// 题目数据为空
	if pd.Data == "" {
		return model.Problem{}, nil
	}

	log.Println("请求解析题目：" + pd.Data)

	// 请求模型
	resp, err := open_ai.Chat(prompt.ParseSystem.String(), pd.Data)
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
