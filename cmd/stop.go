package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop man-hour robot.",
	Long:  `stop man-hour robot.`,
	Run: func(cmd *cobra.Command, args []string) {
		strb, _ := ioutil.ReadFile(".pid.lock")
		command := exec.Command("kill", string(strb))
		command.Start()
		println("service stopped...")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
