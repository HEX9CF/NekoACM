package service

import (
	"encoding/json"
	"errors"
	"log"
	"nekoacm-server/internal/application/dto"
	"nekoacm-server/internal/infrastructure/open_ai"
	"nekoacm-server/pkg/utils"
	"nekoacm-server/prompt"
	"strconv"
	"time"
)

// 生成测试用例
func TestcaseGenerate(ti dto.TestcaseInstruction) (dto.Testcase, error) {
	var t dto.Testcase

	// 说明转换为字符串
	instruction, err := utils.PrettyStruct(ti)
	if err != nil {
		return dto.Testcase{}, err
	}
	log.Println("请求生成测试用例：" + instruction)

	// 请求模型
	resp, err := open_ai.Chat(prompt.TestcaseGenerate, instruction)
	if err != nil {
		log.Println(err)
		return dto.Testcase{}, errors.New("请求模型失败！")
	}
	log.Println("生成结果：" + resp)

	// 解析结果
	err = json.Unmarshal([]byte(resp), &t)
	if err != nil {
		log.Println(err)
		return dto.Testcase{}, errors.New("解析结果失败，请重试！")
	}

	return t, nil
}

// 保存到json文件
func TestcaseSaveJson(t dto.Testcase) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/testcase/" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(t, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
