package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/chat"
	"neko-acm/prompt"
	"os"
	"strings"
)

// 对话
var ChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with AI.",
	Long:  "Chat with AI.",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(" ----- NekoACM Chat ----- ")
		fmt.Println(" exit: 退出")
		fmt.Println(" ------------------------ ")
		fmt.Println(prompt.ChatSystem.Initialization)

		for {
			if err := clearBuffer(reader); err != nil {
				return err
			}
			fmt.Print("> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			op := strings.ToLower(input)
			if op == "exit" {
				break
			}

			msg := model.ChatMsg{
				Content: input,
			}
			resp, err := chat.Assistant(msg)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(resp)
		}

		return nil
	},
}
