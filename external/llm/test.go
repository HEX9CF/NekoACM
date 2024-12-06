package llm

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"stuoj-ai/internal/model"
)

func Test() {
	resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func TestProblemDraft() {
	pi := model.ProblemInstruction{
		Description: "陶陶家的院子里有一棵苹果树，每到秋天树上就会结出 $10$ 个苹果。苹果成熟的时候，陶陶就会跑去摘苹果。陶陶有个 $30$ 厘米高的板凳，当她不能直接用手摘到苹果的时候，就会踩到板凳上再试试。\n\n\n现在已知 $10$ 个苹果到地面的高度，以及陶陶把手伸直的时候能够达到的最大高度，请帮陶陶算一下她能够摘到的苹果的数目。假设她碰到苹果，苹果就会掉下来。",
		Tags:        []string{"模拟"},
	}

	r, err := RequestProblemDraft(pi)
	if err != nil {
		log.Println("Request question failed!")
		return
	}
	log.Println(r)
}
