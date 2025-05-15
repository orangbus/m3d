package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Args:  cobra.NoArgs, // 不允许传参
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
