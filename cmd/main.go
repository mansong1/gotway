package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotway",
	Short: "Simple HTTP API Gateway powered with in-redis cache",
	Long: `Simple HTTP API Gateway powered with in-redis cache.
					Complete documentation available at https://github.com/gotway/gotway`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
