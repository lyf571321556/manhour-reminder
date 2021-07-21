package cmd

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/bot"
	"github.com/lyf571321556/manhour-reminder/conf"
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
		if err := bot.InitBot(viper.ConfigFileUsed()); err != nil {
			log.Fatal(err.Error())
			return
		}
		user := viper.GetString("account")
		password := viper.GetString("password")
		err := testServer(user, password)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("消息发送成功，请查看企业微信.")
	},
}

func testServer(user string, password string) (err error) {
	//获取token
	loginUrl := fmt.Sprintf("%s%s", conf.AppConfig.OnesProjectUrl, service.AUTH_LOGIN)
	_, err = service.Login(loginUrl, user, password)
	if err != nil {
		return err
	}

	testRobotToRemindedUers := make(map[string][]service.UserInfo, 0)
	for _, robot := range conf.AppConfig.RobotList {
		usersMapping := robot.UserMappings
		if _, ok := testRobotToRemindedUers[robot.RobotKey]; !ok {
			testRobotToRemindedUers[robot.RobotKey] = make([]service.UserInfo,
				0)
		}
		botUserList := testRobotToRemindedUers[robot.RobotKey]
		for _, userMapping := range usersMapping {
			var user = service.UserInfo{
				UUID:       userMapping.OnesUserid,
				WechatUUID: userMapping.WechatUserid,
			}
			botUserList = append(botUserList, user)
		}
		testRobotToRemindedUers[robot.RobotKey] = botUserList

	}

	err = bot.SendMsgToUser(testRobotToRemindedUers)
	return err
}