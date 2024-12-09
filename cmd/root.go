package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "neko",
	Short: "A large model-based ACM-ICPC algorithm problem automatic generation system. ",
	Long:  "A large model-based ACM-ICPC algorithm problem automatic generation system that can automatically generate algorithm problems, testcases, and problem solutions. ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run neko-acm-ai")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
