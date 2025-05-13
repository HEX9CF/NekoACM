package service

import (
	"encoding/json"
	"errors"
	"log"
	"nekoacm-server/internal/application/dto"
	"nekoacm-server/internal/infrastructure/openai"
	"nekoacm-server/pkg/utils"
	"nekoacm-server/prompt"
)

// 提交评测
func Submit(s dto.Submission) (dto.Judgement, error) {
	var j dto.Judgement

	// 转换为字符串
	submission, err := utils.PrettyStruct(s)
	if err != nil {
		return dto.Judgement{}, err
	}
	log.Println("请求评测：" + submission)

	// 请求模型
	resp, err := openai.Chat(prompt.JudgeSubmit, submission)
	if err != nil {
		log.Println(err)
		return dto.Judgement{}, errors.New("请求模型失败！")
	}
	log.Println("评测结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &j)
	if err != nil {
		log.Println(err)
		return dto.Judgement{}, errors.New("解析结果失败，请重试！")
	}

	return j, nil
}
