package main

import (
	"github.com/gotway/gotway/cmd/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server",
	Long:  "HTTP API gateway and Kubernetes controller",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
