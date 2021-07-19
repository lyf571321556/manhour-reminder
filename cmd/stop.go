package cmd

import (
	"github.com/lyf571321556/manhour-reminder/log"
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
		log.Info("service stopped...")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
