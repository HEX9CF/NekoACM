package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/solution"
	"os"
	"strings"
)

// 生成题解
var SolutionCmd = &cobra.Command{
	Use:   "solution",
	Short: "Generate a solution",
	Long:  "Generate a solution.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 生成题解 -------- ")
		reader := bufio.NewReader(os.Stdin)
		si, err := readSolutionInstruction(reader)
		if err != nil {
			return err
		}

		for {
			// 生成题解
			fmt.Println("正在生成题解...")
			s, err := solution.Generate(si)
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
			fmt.Println("题解生成成功")

			// 保存到文件
			err = saveSolutionJson(reader, s)
			if err != nil {
				log.Println(err)
			}

			// 继续生成
			_, err = reader.Discard(reader.Buffered())
			if err != nil {
				return err
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

// 读取题解信息
func readSolutionInstruction(reader *bufio.Reader) (model.SolutionInstruction, error) {
	si := model.SolutionInstruction{}
	err := clearBuffer(reader)
	if err != nil {
		return model.SolutionInstruction{}, err
	}

	fmt.Println("请输入题目信息：")
	fmt.Print("标题：")
	si.Title, _ = reader.ReadString('\n')
	si.Title = strings.TrimSpace(si.Title)
	fmt.Print("描述：")
	si.Description, _ = reader.ReadString('\n')
	si.Description = strings.TrimSpace(si.Description)
	fmt.Print("输入说明：")
	si.Input, _ = reader.ReadString('\n')
	si.Input = strings.TrimSpace(si.Input)
	fmt.Print("输出说明：")
	si.Output, _ = reader.ReadString('\n')
	si.Output = strings.TrimSpace(si.Output)
	fmt.Print("样例输入：")
	si.SampleInput, _ = reader.ReadString('\n')
	si.SampleInput = strings.TrimSpace(si.SampleInput)
	fmt.Print("样例输出：")
	si.SampleOutput, _ = reader.ReadString('\n')
	si.SampleOutput = strings.TrimSpace(si.SampleOutput)
	fmt.Print("提示：")
	si.Hint, _ = reader.ReadString('\n')
	si.Hint = strings.TrimSpace(si.Hint)
	fmt.Print("标签（以空格分隔）：")
	tagsInput, _ := reader.ReadString('\n')
	si.Tags = strings.Fields(tagsInput)
	fmt.Print("已有题解代码：")
	si.Solution, _ = reader.ReadString('\n')
	si.Solution = strings.TrimSpace(si.Solution)
	fmt.Print("目标编程语言：")
	si.Language, _ = reader.ReadString('\n')
	si.Language = strings.TrimSpace(si.Language)

	return si, nil
}

// 保存题解到文件
func saveSolutionJson(reader *bufio.Reader, s model.Solution) error {
	err := clearBuffer(reader)
	if err != nil {
		return err
	}
	fmt.Print("是否保存到文件(Y/N)?")
	save, _ := reader.ReadString('\n')
	save = strings.TrimSpace(save)
	save = strings.ToLower(save)

	if save == "y" {
		path, err := solution.SaveJson(s)
		if err != nil {
			fmt.Println("保存失败！")
			return err
		}
		fmt.Println("保存成功，文件路径：" + path)
	}

	return nil
}
