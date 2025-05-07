package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"neko-acm/cmd/bootstrap"
)

// 服务器命令
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the server.",
	Long:  "Run the server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("正在启动服务器...")
		return bootstrap.InitServer()
	},
}
