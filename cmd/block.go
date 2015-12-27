package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Block a command from running.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkBlockFlags()
	},
	Long: `Add a block to the KV store so that the command will not execute.`,
	Run:  startBlock,
}

func startBlock(cmd *cobra.Command, args []string) {
	status := BlockCommand()
	if status {
		Log(fmt.Sprintf("'%s' was blocked.", Exec), "info")
		StatsdSend("block")
	} else {
		Log(fmt.Sprintf("'%s' was NOT blocked - something went wrong.", Exec), "info")
	}
}

func checkBlockFlags() {
	Log("block: Checking cli flags.", "debug")
	if Exec == "" {
		fmt.Println("Need a command to block with '-e'")
		os.Exit(0)
	}
	Log("block: Required cli flags present.", "debug")
}

var (
	// Reason to block command (optional)
	Reason string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&Reason, "reason", "r", "", "Reason to block command.")
	RootCmd.AddCommand(blockCmd)
}
