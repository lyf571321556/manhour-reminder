package cmd

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/conf"
	"github.com/lyf571321556/manhour-reminder/robot"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	rootCmd.AddCommand(test)
}

var test = &cobra.Command{
	Use:   "test",
	Short: "take a test for Robot",
	Long:  "take a test for Robot,and send manhour reminded message to all user in config's user mappings.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := robot.InitBot(viper.ConfigFileUsed()); err != nil {
			log.Fatal(err.Error())
			return
		}
		err := testServer()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("消息发送成功，请查看企业微信.")
	},
}

func testServer() (err error) {
	user := viper.GetString("account")
	password := viper.GetString("password")
	loginUrl := fmt.Sprintf("%s%s", conf.AppConfig.OnesProjectUrl, service.AUTH_LOGIN)
	AppAuth, err := service.Login(loginUrl, user, password)
	if err != nil {
		return err
	}
	err = robot.StartCheckUsersManhourInEveryRobot(AppAuth)
	return err
}
