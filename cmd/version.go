package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of filetool",
	Long:  "Print the version of Filetool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Filetool V0.0.1")
	},
}
