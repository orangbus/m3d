/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/orangbus/m3d/pkg/github"
	"github.com/spf13/cobra"
	"log"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "测试方法",
	Run: func(cmd *cobra.Command, args []string) {
		//str := "156b1176eca78e4e63cc96bb920e627ac72ababa"
		//log.Println(str[:7])
		//return
		g := github.NewGithub()
		//res, err := g.Get()
		//if err != nil {
		//	log.Print(err.Error())
		//	return
		//}
		//log.Println(res)

		if err := g.Download(); err != nil {
			log.Print(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
