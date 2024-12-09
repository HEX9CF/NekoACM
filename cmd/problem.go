package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"neko-acm/internal/model"
	"neko-acm/internal/service/problem"
	"neko-acm/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

var ProblemCmd = &cobra.Command{
	Use:   "problem",
	Short: "Generate a problem",
	Long:  "Generate a problem.",
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
		problem, err := problem.Draft(pi)
		if err != nil {
			return err
		}
		fmt.Println("题目生成成功！")

		fmt.Print("是否保存到文件(Y/N)?")
		save, _ := reader.ReadString('\n')
		save = strings.TrimSpace(save)
		save = strings.ToLower(save)

		if save == "y" {
			timestamp := time.Now().Unix()
			path := "output/problem/" + problem.Title + "_" + strconv.FormatInt(timestamp, 10) + ".json"
			err := utils.WriteJson(problem, path)
			if err != nil {
				fmt.Println("保存失败！")
				return err
			} else {
				fmt.Println("保存成功，文件路径：" + path)
			}
		}
		return nil
	},
}
