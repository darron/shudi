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

var (
	// Delay is the maximum amount of time to delay execution of Exec.
	Delay int

	// Exec is the command that should be executed if there's no block.
	Exec string

	// Prefix is the location in Consul's KV store to keep state information.
	Prefix string

	// Splay is the maximum amount of time we should deduct
	// from Delay for the first execution of Exec.
	Splay int

	// Token is used for access to Consul if an ACL is being used.
	Token string

	// Verbose logs all output to stdout.
	Verbose bool
)

func init() {
	RootCmd.PersistentFlags().IntVarP(&Delay, "delay", "d", 60, "maximum amount of time to delay execution")
	RootCmd.PersistentFlags().IntVarP(&Splay, "splay", "s", 30, "maximum amount of time to splay execution")
	RootCmd.PersistentFlags().StringVarP(&Exec, "exec", "e", "", "Execute this command if there's no block.")
	RootCmd.PersistentFlags().StringVarP(&Prefix, "prefix", "p", "shudi", "Consul prefix for saved state.")
	RootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "anonymous", "Token for Consul access")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "log output to stdout")
}
