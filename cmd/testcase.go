package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"neko-acm/internal/model"
	"neko-acm/internal/service/testcase"
	"neko-acm/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

var TestcaseCmd = &cobra.Command{
	Use:   "testcase",
	Short: "Generate a testcase",
	Long:  "Generate a testcase.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 生成测试用例 -------- ")
		reader := bufio.NewReader(os.Stdin)
		pi := readTestcaseInstruction(reader)

		for {
			// 生成题目
			fmt.Println("正在生成测试用例...")
			testcase, err := testcase.Draft(pi)
			if err != nil {
				fmt.Println("测试用例生成失败！")
				log.Println(err)
				continue
			}
			fmt.Println("测试用例生成成功")

			// 保存到文件
			err = saveTestcaseJson(reader, testcase)
			if err != nil {
				log.Println(err)
			}

			fmt.Print("是否继续生成测试用例(Y/N)?")
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

func readTestcaseInstruction(reader *bufio.Reader) model.TestcaseInstruction {
	ti := model.TestcaseInstruction{}

	// 读取题目信息
	fmt.Println("请输入题目信息：")
	fmt.Print("标题：")
	ti.Title, _ = reader.ReadString('\n')
	fmt.Print("描述：")
	ti.Description, _ = reader.ReadString('\n')
	fmt.Print("输入说明：")
	ti.Input, _ = reader.ReadString('\n')
	fmt.Print("输出说明：")
	ti.Output, _ = reader.ReadString('\n')
	fmt.Print("样例输入：")
	ti.SampleInput, _ = reader.ReadString('\n')
	fmt.Print("样例输出：")
	ti.SampleOutput, _ = reader.ReadString('\n')
	fmt.Print("提示：")
	ti.Hint, _ = reader.ReadString('\n')
	fmt.Print("标签（以空格分隔）：")
	tagsInput, _ := reader.ReadString('\n')
	ti.Tags = strings.Fields(tagsInput)
	fmt.Print("题解代码：")
	ti.Solution, _ = reader.ReadString('\n')

	return ti
}

func saveTestcaseJson(reader *bufio.Reader, testcase model.Testcase) error {
	fmt.Print("是否保存到文件(Y/N)?")
	save, _ := reader.ReadString('\n')
	save = strings.TrimSpace(save)
	save = strings.ToLower(save)

	if save == "y" {
		timestamp := time.Now().Unix()
		path := "output/testcase/" + strconv.FormatInt(timestamp, 10) + ".json"
		err := utils.WriteJson(testcase, path)
		if err != nil {
			fmt.Println("保存失败！")
			return err
		}
		fmt.Println("保存成功，文件路径：" + path)
	}

	return nil
}
