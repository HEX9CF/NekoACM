package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/problem"
	"os"
	"strings"
)

// 翻译题目
var TranslateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate a problem",
	Long:  "Translate an algorithm problem to target language.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 翻译题目 -------- ")
		reader := bufio.NewReader(os.Stdin)
		ti, err := readTranslateInstruction(reader)
		if err != nil {
			return err
		}

		for {
			// 翻译题目
			fmt.Println("正在翻译题目...")
			p, err := problem.Translate(ti)
			if err != nil {
				log.Println(err)

				// 重试
				err := clearBuffer(reader)
				if err != nil {
					return err
				}
				fmt.Print("翻译失败，是否重试(Y/N)?")
				again, _ := reader.ReadString('\n')
				again = strings.TrimSpace(again)
				again = strings.ToLower(again)
				if again != "y" {
					break
				}
				continue
			}
			fmt.Println("题目翻译成功")

			// 保存到文件
			err = saveProblemJson(reader, p)
			if err != nil {
				log.Println(err)
			}

			// 重新翻译
			_, err = reader.Discard(reader.Buffered())
			if err != nil {
				return err
			}
			fmt.Print("是否重新翻译题目(Y/N)?")
			again, _ := reader.ReadString('\n')
			again = strings.TrimSpace(again)
			again = strings.ToLower(again)
			if again != "y" {
				break
			} else {
				ti, err = readTargetLang(reader, ti)
				if err != nil {
					return err
				}
			}
		}

		return nil
	},
}

// 读取翻译题目信息
func readTranslateInstruction(reader *bufio.Reader) (model.TranslateInstruction, error) {
	ti := model.TranslateInstruction{}
	err := clearBuffer(reader)
	if err != nil {
		return model.TranslateInstruction{}, err
	}

	fmt.Println("请输入题目信息：")
	fmt.Print("标题：")
	ti.Title, _ = reader.ReadString('\n')
	ti.Title = strings.TrimSpace(ti.Title)
	fmt.Print("描述：")
	ti.Description, _ = reader.ReadString('\n')
	ti.Description = strings.TrimSpace(ti.Description)
	fmt.Print("输入说明：")
	ti.Input, _ = reader.ReadString('\n')
	ti.Input = strings.TrimSpace(ti.Input)
	fmt.Print("输出说明：")
	ti.Output, _ = reader.ReadString('\n')
	ti.Output = strings.TrimSpace(ti.Output)
	fmt.Print("样例输入：")
	ti.SampleInput, _ = reader.ReadString('\n')
	ti.SampleInput = strings.TrimSpace(ti.SampleInput)
	fmt.Print("样例输出：")
	ti.SampleOutput, _ = reader.ReadString('\n')
	ti.SampleOutput = strings.TrimSpace(ti.SampleOutput)
	fmt.Print("提示：")
	ti.Hint, _ = reader.ReadString('\n')
	ti.Hint = strings.TrimSpace(ti.Hint)
	fmt.Print("标签（以空格分隔）：")
	tagsInput, _ := reader.ReadString('\n')
	ti.Tags = strings.Fields(tagsInput)
	ti, err = readTargetLang(reader, ti)
	if err != nil {
		return model.TranslateInstruction{}, err
	}

	return ti, nil
}

func readTargetLang(reader *bufio.Reader, ti model.TranslateInstruction) (model.TranslateInstruction, error) {
	fmt.Print("目标语言：")
	ti.TargetLang, _ = reader.ReadString('\n')
	ti.TargetLang = strings.TrimSpace(ti.TargetLang)

	return ti, nil
}
