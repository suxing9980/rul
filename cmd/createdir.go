package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"rul/utils/gout"
)

var (
	prefix    string
	total     int
	targetDir string
)

var createdirCmd = &cobra.Command{
	Use:   "createdir",
	Short: "createDirs",
	Long:  `Create directories with a prefix and a range of numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		createDirs(prefix, total, targetDir)
	},
}

func init() {
	rootCmd.AddCommand(createdirCmd)
	createdirCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Prefix for the directories")
	createdirCmd.Flags().IntVarP(&total, "total", "t", 0, "Total number of directories to create")
	createdirCmd.Flags().StringVarP(&targetDir, "target", "d", "", "Target directory to create the directories in")
}

func createDirs(prefix string, total int, targetDir string) {
	for i := 1; i <= total; i++ {
		dirName := fmt.Sprintf("%s%d", prefix, i)
		dirPath := filepath.Join(targetDir, dirName)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			gout.RedSPrintf("Error creating directory %s: %v\n\n", dirPath, err)
		} else {
			gout.GreenSPrintf("Created directory %s\n", dirPath)
		}
	}
}
