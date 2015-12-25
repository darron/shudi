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

	// Backend is where the saved state is stored.
	Backend string

	// Prefix is the location in the KV store to keep state information.
	Prefix string

	// Connection is how to connect to the KV store.
	Connection string

	// Splay is the maximum amount of time we should deduct
	// from Delay for the first execution of Exec.
	Splay int

	// Token is used for access to Consul if an ACL is being used.
	Token string

	// Verbose logs all output to stdout.
	Verbose bool

	// Once only runs once - and doesn't loop.
	Once bool

	// DogStatsd sends metrics to the local Dogstatsd endpoint.
	DogStatsd bool
)

func init() {
	AutoEnable()
	RootCmd.PersistentFlags().IntVarP(&Delay, "delay", "d", 60, "maximum amount of time to delay execution")
	RootCmd.PersistentFlags().IntVarP(&Splay, "splay", "s", 30, "maximum amount of time to splay execution")
	RootCmd.PersistentFlags().StringVarP(&Exec, "exec", "e", "", "Execute this command if there's no block.")
	RootCmd.PersistentFlags().StringVarP(&Backend, "backend", "b", "consul", "Backend for saved state.")
	RootCmd.PersistentFlags().StringVarP(&Prefix, "prefix", "p", "shudi", "Prefix for saved state.")
	RootCmd.PersistentFlags().StringVarP(&Connection, "connection", "c", "127.0.0.1:8500", "Connection string for saved state.")
	RootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "anonymous", "Token for Consul access")
	RootCmd.PersistentFlags().BoolVarP(&Once, "once", "o", false, "Don't loop - just run once.")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "log output to stdout")
	RootCmd.PersistentFlags().BoolVarP(&DogStatsd, "dogstatsd", "", false, "Send metrics to Dogstatsd.")
}
