package cli

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/service/misc"
	"os"
	"strings"
)

// 生成笑话
var JokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Tell a joke.",
	Long:  "Tell a joke about algorithm competition to make you happy.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 生成笑话 -------- ")
		reader := bufio.NewReader(os.Stdin)

		for {
			// 生成笑话
			fmt.Println("正在生成笑话...")
			_, err := misc.TellJoke()
			if err != nil {
				log.Println(err)

				// 重试
				err := clearBuffer(reader)
				if err != nil {
					return err
				}
				fmt.Print("生成失败，是否重试(Y/N)?")
				again, _ := reader.ReadString('\n')
				again = strings.TrimSpace(again)
				again = strings.ToLower(again)
				if again != "y" {
					break
				}
				continue
			}
			fmt.Println("笑话生成成功")

			// 继续生成
			_, err = reader.Discard(reader.Buffered())
			if err != nil {
				return err
			}
			fmt.Print("要不要再来一个(Y/N)?")
			again, _ := reader.ReadString('\n')
			again = strings.TrimSpace(again)
			again = strings.ToLower(again)
			if again != "y" {
				break
			}
		}

		return nil
	},
}
