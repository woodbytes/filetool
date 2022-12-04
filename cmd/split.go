package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(splitCmd)
}

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a file ",
	Long:  "Split a file in the number of ...",
	Run:   splitRun,
}

func splitRun(cmd *cobra.Command, args []string) {
	fmt.Println("arg0: ", args[0])
	fmt.Println("arg1: ", args[1])
}
