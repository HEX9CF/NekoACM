package judge

import (
	"encoding/json"
	"errors"
	"log"
	"neko-acm/internal/infrastructure/open_ai"
	"neko-acm/internal/model"
	"neko-acm/pkg/utils"
	"neko-acm/prompt"
)

// 提交评测
func Submit(s model.Submission) (model.Judgement, error) {
	var j model.Judgement

	// 转换为字符串
	submission, err := utils.PrettyStruct(s)
	if err != nil {
		return model.Judgement{}, err
	}
	log.Println("请求评测：" + submission)

	// 请求模型
	resp, err := open_ai.Chat(prompt.JudgeSubmit, submission)
	if err != nil {
		log.Println(err)
		return model.Judgement{}, errors.New("请求模型失败！")
	}
	log.Println("评测结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &j)
	if err != nil {
		log.Println(err)
		return model.Judgement{}, errors.New("解析结果失败，请重试！")
	}

	return j, nil
}
