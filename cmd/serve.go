/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/bootstrap"
	"github.com/orangbus/m3d/pkg/config"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "web服务器",
	Long:  `启动一个web服务器， m3d serve 访问端口号：3000`,
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.New()
		bootstrap.SetupRouter(router)

		err := router.Run(fmt.Sprintf(":%d", config.GetInt("app.port")))
		if err != nil {
			log.Printf("服务启动失败:%s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
