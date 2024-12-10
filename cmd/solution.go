package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/solution"
	"neko-acm/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

var SolutionCmd = &cobra.Command{
	Use:   "solution",
	Short: "Generate a solution",
	Long:  "Generate a solution.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 生成题解 -------- ")
		reader := bufio.NewReader(os.Stdin)
		pi := readSolutionInstruction(reader)

		for {
			// 生成题目
			fmt.Println("正在生成题解...")
			solution, err := solution.Draft(pi)
			if err != nil {
				fmt.Println("题解生成失败！")
				log.Println(err)
				continue
			}
			fmt.Println("题解生成成功")

			// 保存到文件
			err = saveSolutionJson(reader, solution)
			if err != nil {
				log.Println(err)
			}

			fmt.Print("是否继续生成题解(Y/N)?")
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

func readSolutionInstruction(reader *bufio.Reader) model.SolutionInstruction {
	si := model.SolutionInstruction{}

	// 读取题目信息
	fmt.Println("请输入题目信息：")
	fmt.Print("标题：")
	si.Title, _ = reader.ReadString('\n')
	fmt.Print("描述：")
	si.Description, _ = reader.ReadString('\n')
	fmt.Print("输入说明：")
	si.Input, _ = reader.ReadString('\n')
	fmt.Print("输出说明：")
	si.Output, _ = reader.ReadString('\n')
	fmt.Print("样例输入：")
	si.SampleInput, _ = reader.ReadString('\n')
	fmt.Print("样例输出：")
	si.SampleOutput, _ = reader.ReadString('\n')
	fmt.Print("提示：")
	si.Hint, _ = reader.ReadString('\n')
	fmt.Print("标签（以空格分隔）：")
	tagsInput, _ := reader.ReadString('\n')
	si.Tags = strings.Fields(tagsInput)
	fmt.Print("已有题解代码：")
	si.Solution, _ = reader.ReadString('\n')
	fmt.Print("目标编程语言：")
	si.Language, _ = reader.ReadString('\n')

	return si
}

func saveSolutionJson(reader *bufio.Reader, solution model.Solution) error {
	fmt.Print("是否保存到文件(Y/N)?")
	save, _ := reader.ReadString('\n')
	save = strings.TrimSpace(save)
	save = strings.ToLower(save)

	if save == "y" {
		timestamp := time.Now().Unix()
		path := "output/solution/" + solution.Language + "_" + strconv.FormatInt(timestamp, 10) + ".json"
		err := utils.WriteJson(solution, path)
		if err != nil {
			fmt.Println("保存失败！")
			return err
		}
		fmt.Println("保存成功，文件路径：" + path)
	}

	return nil
}
