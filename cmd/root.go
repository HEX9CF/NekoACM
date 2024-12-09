package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "neko",
	Short: "A large model-based ACM-ICPC algorithm problem automatic generation system",
	Long:  "A large model-based ACM-ICPC algorithm problem automatic generation system that can automatically generate algorithm problems, testcases, and problem solutions. ",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			//fmt.Println(input)
			switch input {
			case "exit":
				os.Exit(0)
			case "server":
				if err := ServerCmd.RunE(nil, nil); err != nil {
					return err
				}
			case "problem":
				if err := ProblemCmd.RunE(nil, nil); err != nil {
					return err
				}
			default:
				fmt.Println("未知命令！")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ServerCmd)
	rootCmd.AddCommand(ProblemCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
