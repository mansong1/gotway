package main

import (
	"github.com/gotway/gotway/cmd/cache"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache",
	Long:  "Cache utilities",
	Args:  cobra.NoArgs,
}

var cacheInvalidateCmd = &cobra.Command{
	Use:   "invalidate",
	Short: "Invalidate cache",
	Long:  "Invalidate cache",
	Args:  cobra.NoArgs,
}

var tags []string

var cacheInvalidateTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Invalidate cache by tags",
	Long:  "Invalidate cache by tags",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cache.InvalidateByTags(tags)
	},
}

var service string
var paths []string

var cacheInvalidatePathsCmd = &cobra.Command{
	Use:   "invalidate",
	Short: "Invalidate cache by paths",
	Long:  "Invalidate cache by paths",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cache.InvalidateByPaths(service, paths)
	},
}

func init() {
	flags := cacheInvalidateCmd.PersistentFlags()

	flags.StringSliceVarP(&tags, "tags", "t", []string{}, "Tags to invalidate")
	flags.StringVarP(&service, "service", "s", "", "Service cache to invalidate")
	flags.StringSliceVarP(&paths, "paths", "p", []string{}, "Paths to invalidate")

	cacheInvalidateCmd.AddCommand(cacheInvalidatePathsCmd)
	cacheInvalidateCmd.AddCommand(cacheInvalidateTagsCmd)
	cacheCmd.AddCommand(cacheInvalidateCmd)
	rootCmd.AddCommand(cacheCmd)
}
