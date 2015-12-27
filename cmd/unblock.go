package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Unblock a command from running.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkUnblockFlags()
	},
	Long: `Remove a block to the KV store so that the command will execute.`,
	Run:  startUnblock,
}

func startUnblock(cmd *cobra.Command, args []string) {
	status := UnblockCommand()
	if status {
		Log(fmt.Sprintf("'%s' was unblocked.", Exec), "info")
		StatsdSend("unblock")
	} else {
		Log(fmt.Sprintf("'%s' was NOT unblocked - something went wrong.", Exec), "info")
	}
}

func checkUnblockFlags() {
	Log("unblock: Checking cli flags.", "debug")
	if Exec == "" {
		fmt.Println("Need a command to unblock with '-e'")
		os.Exit(0)
	}
	Log("unblock: Required cli flags present.", "debug")
}

var ()

func init() {
	RootCmd.AddCommand(unblockCmd)
}
