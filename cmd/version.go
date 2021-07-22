package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version string
var Time string
var User string

func init() {
	rootCmd.AddCommand(showVersion)
}

var showVersion = &cobra.Command{
	Use:   "version",
	Short: "show version for manhour-robot",
	Long:  "show version for manhour-robot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build Time:\t", Time)
		fmt.Println("Build User:\t", User)
		fmt.Println("Build Version:\t", version)
	},
}
