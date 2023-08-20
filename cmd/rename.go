package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"rul/utils/gout"
	"strings"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "文件改名工具",
	Long:  `把一个指定目录下所有文件的名称去掉给出的字符串`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		str, _ := cmd.Flags().GetString("str")
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				gout.RedSPrintf("出现错误%s\n", err.Error())
				return err
			}
			newName := strings.ReplaceAll(info.Name(), str, "")
			if newName != info.Name() {
				os.Rename(path, filepath.Join(filepath.Dir(path), newName))
			}
			if info.IsDir() {
				gout.GreenSPrintf("目录:%s\trename成功", info.Name())
			} else {
				gout.GreenSPrintf("文件:%s\trename成功", info.Name())
			}

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.Flags().StringP("dir", "d", ".", "指定目标dir")
	renameCmd.Flags().StringP("str", "s", "", "指定需要移除的str")
}
