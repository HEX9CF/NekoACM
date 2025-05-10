package cli

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/application/dto"
	"neko-acm/internal/application/service"
	"os"
	"strings"
)

// 解析题目
var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a problem.",
	Long:  "Parse a problem from other format to JSON.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 解析题目 -------- ")
		reader := bufio.NewReader(os.Stdin)
		pd := dto.ProblemData{}
		if err := clearBuffer(reader); err != nil {
			return err
		}
		fmt.Println("请输入题目数据：")
		pd.Data, _ = reader.ReadString('\n')
		pd.Data = strings.TrimSpace(pd.Data)

		for {

			// 生成题目
			fmt.Println("正在解析题目...")
			p, err := service.ProblemParse(pd)
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
			fmt.Println("题目解析成功")

			// 保存到文件
			err = saveProblemJson(reader, p)
			if err != nil {
				log.Println(err)
			}
		}

		return nil
	},
}
