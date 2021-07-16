package cmd

import (
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop man-hour rebot.",
	Long:  `stop man-hour rebot.`,
	Run: func(cmd *cobra.Command, args []string) {
		//strb, _ := ioutil.ReadFile(".pid.lock")
		//command := exec.Command("kill", string(strb))
		//command.Start()
		println("service stop...")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
