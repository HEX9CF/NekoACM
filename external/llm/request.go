package llm

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"log"
	"stuoj-ai/internal/model"
	"stuoj-ai/utils"
)

func Request(msg openai.ChatCompletionMessage) (openai.ChatCompletionMessage, error) {
	resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    config.Model,
			Messages: []openai.ChatCompletionMessage{msg},
		},
	)

	if err != nil {
		return openai.ChatCompletionMessage{}, errors.New("ChatCompletion error: " + err.Error())
	}

	return resp.Choices[0].Message, nil
}

func RequestProblemDraft(pi model.ProblemInstruction) (string, error) {
	instruction, err := utils.PrettyStruct(pi)
	log.Println(instruction)
	if err != nil {
		return "", err
	}

	resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: JsonInputPrompt +
						"\n\n" + JsonOutputPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: instruction,
				},
			},
		},
	)

	if err != nil {
		return "", errors.New("ChatCompletion error: " + err.Error())
	}

	return resp.Choices[0].Message.Content, nil
}
