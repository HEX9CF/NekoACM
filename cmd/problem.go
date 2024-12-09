package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"neko-acm/internal/model"
	"neko-acm/internal/service/problem"
	"os"
	"strings"
)

var ProblemCmd = &cobra.Command{
	Use:   "problem",
	Short: "Run the server.",
	Long:  "Run the server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		pi := model.ProblemInstruction{}

		fmt.Println("请输入题目信息：")
		fmt.Print("标题：")
		pi.Title, _ = reader.ReadString('\n')
		fmt.Print("描述：")
		pi.Description, _ = reader.ReadString('\n')
		fmt.Print("输入说明：")
		pi.Input, _ = reader.ReadString('\n')
		fmt.Print("输出说明：")
		pi.Output, _ = reader.ReadString('\n')
		fmt.Print("样例输入：")
		pi.SampleInput, _ = reader.ReadString('\n')
		fmt.Print("样例输出：")
		pi.SampleOutput, _ = reader.ReadString('\n')
		fmt.Print("提示：")
		pi.Hint, _ = reader.ReadString('\n')
		fmt.Print("标签（以空格分隔）：")
		tagsInput, _ := reader.ReadString('\n')
		pi.Tags = strings.Fields(tagsInput)
		fmt.Print("题解代码：")
		pi.Solution, _ = reader.ReadString('\n')

		fmt.Println("正在生成题目...")
		_, err := problem.Draft(pi)
		if err != nil {
			return err
		}
		fmt.Println("题目生成成功！")
		return nil
	},
}
