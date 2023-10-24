package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"rul/utils/gout"
)

var rootCmd = &cobra.Command{
	Use:   "rul",
	Short: "rul工具",
	Long:  `常用命令行工具`,
	Run: func(cmd *cobra.Command, args []string) {
		gout.GreenSPrintf("\t1、rename 用于对指定dir目录所有文件去除固定字符串\t%s", `.\rul.exe rename -d ".\a.b" -s "as.*bn"`)
		gout.GreenSPrintf("\t2、createdir 用于对指定dir目录创建带有统一前缀的目录\t%s", `.\rul.exe createdir -p "chapter" -d "./" -t 10`)
		return
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
