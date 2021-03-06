package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	waitTime := GetTime()
	Sleep(waitTime)
	for {
		if CheckForBlock() {
			RunCommand(Exec)
			StatsdSend("run")
		} else {
			Log(fmt.Sprintf("skip='true' noexec='%s'", Exec), "info")
			StatsdSend("skip")
		}
		if Once {
			os.Exit(0)
		} else {
			Sleep(Delay)
		}
	}
}

func checkRunFlags() {
	Log("run: Checking cli flags.", "debug")
	if Exec == "" {
		fmt.Println("Need a command to exec with '-e'")
		os.Exit(0)
	}
	Log("run: Required cli flags present.", "debug")
}

var ()

func init() {
	RootCmd.AddCommand(runCmd)
}
