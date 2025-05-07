package solution

import (
	"encoding/json"
	"errors"
	"log"
	"neko-acm/internal/infrastructure/open_ai"
	"neko-acm/internal/model"
	"neko-acm/prompt"
	"neko-acm/utils"
)

// 生成题解
func Generate(si model.SolutionInstruction) (model.Solution, error) {
	var s model.Solution

	// 说明转换为字符串
	instruction, err := utils.PrettyStruct(si)
	if err != nil {
		return model.Solution{}, err
	}
	log.Println("请求生成题解：" + instruction)

	// 请求模型
	resp, err := open_ai.Chat(prompt.SolutionGenerate, instruction)
	if err != nil {
		log.Println(err)
		return model.Solution{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		log.Println(err)
		return model.Solution{}, errors.New("解析结果失败，请重试！")
	}

	return s, nil
}
