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

// 生成题目
func ProblemGenerate(pi dto.ProblemInstruction) (dto.Problem, error) {
	var p dto.Problem

	// 题目说明转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return dto.Problem{}, err
	}
	log.Println("请求生成题目：" + instruction)

	// 请求模型
	resp, err := openai.Chat(prompt.ProblemGenerate, instruction)
	if err != nil {
		log.Println(err)
		return dto.Problem{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		log.Println(err)
		return dto.Problem{}, errors.New("解析结果失败，请重试！")
	}

	return p, nil
}

// 解析题目
func ProblemParse(pd dto.ProblemData) (dto.Problem, error) {
	var p dto.Problem

	// 题目数据为空
	if pd.Data == "" {
		return dto.Problem{}, nil
	}

	log.Println("请求解析题目：" + pd.Data)

	// 请求模型
	resp, err := openai.Chat(prompt.ProblemParse, pd.Data)
	if err != nil {
		return dto.Problem{}, errors.New("请求模型失败！")
	}
	log.Println("解析结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		log.Println(err)
		return dto.Problem{}, errors.New("解析结果失败，请重试！")
	}

	return p, nil
}

// 翻译题目
func ProblemTranslate(pi dto.TranslateInstruction) (dto.Problem, error) {
	var p dto.Problem

	// 题目说明转换为字符串
	instruction, err := utils.PrettyStruct(pi)
	if err != nil {
		return dto.Problem{}, err
	}
	log.Println("请求翻译题目：" + instruction)

	// 请求模型
	resp, err := openai.Chat(prompt.ProblemTranslate, instruction)
	if err != nil {
		return dto.Problem{}, errors.New("请求模型失败！")
	}
	log.Println("翻译结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &p)
	if err != nil {
		log.Println(err)
		return dto.Problem{}, errors.New("解析结果失败，请重试！")
	}

	return p, nil
}

// 保存到json文件
func ProblemSaveJson(p dto.Problem) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/problem/" + p.Title + "_" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(p, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
