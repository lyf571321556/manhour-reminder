package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version string

func init() {
	rootCmd.AddCommand(version)
}

var version = &cobra.Command{
	Use:   "version",
	Short: "show version for manhour-robot",
	Long:  "show version for manhour-robot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
