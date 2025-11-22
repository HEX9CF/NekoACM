package stdio

import (
	"bufio"
	"fmt"
	"log"
	"nekoacm/internal/application/dto"
	"nekoacm/internal/application/service"
	"os"
	"strings"

	"github.com/spf13/cobra"
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

			msg := dto.ChatMsg{
				Content: input,
			}
			resp, err := service.AssistantChat(msg)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(resp)
		}

		return nil
	},
}
