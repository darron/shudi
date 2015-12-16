package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// RootCmd is the base command that sets up all other commands.
var RootCmd = &cobra.Command{
	Use:   "shudi",
	Short: "Should I run?",
	Long:  `Small binary that runs a job if it's not blocked.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("`shudi -h` for help information.")
		fmt.Println("`shudi -v` for version information.")
	},
}

var ()

func init() {
}
