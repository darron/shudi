package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Should I run?",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkRunFlags()
	},
	Long: `Should I run? Check Consul's KV store and run if you're not blocked.`,
	Run:  startRun,
}

func startRun(cmd *cobra.Command, args []string) {
}

func checkRunFlags() {
	Log("run: Checking cli flags.", "debug")

	Log("run: Required cli flags present.", "debug")
}

var ()

func init() {
	RootCmd.AddCommand(runCmd)
}
