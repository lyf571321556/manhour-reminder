package bot

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/conf"
	"github.com/lyf571321556/manhour-reminder/log"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"github.com/lyf571321556/qiye-wechat-bot-api/bot"
	"github.com/lyf571321556/qiye-wechat-bot-api/text"
)

var wechatbot map[string]api.QiyeWechatBot

func InitBot(configPath string) (err error) {
	err = conf.Init(configPath)
	if err != nil {
		return err
	}
	err = log.InitLog()
	if err != nil {
		return err
	}
	wechatbot = make(map[string]api.QiyeWechatBot, 0)
	for _, botInfo := range conf.AppConfig.RobotList {
		bot := bot.NewQiyeWechatBot(botInfo.RobotKey)
		api.SetDebug(true)
		wechatbot[botInfo.RobotKey] = bot
	}
	return err
}

func SendMsgToUser(robotToRemindedUers map[string][]service.UserInfo) (err error) {
	if robotToRemindedUers == nil {
		return err
	}
	for botKey, userList := range robotToRemindedUers {
		robot, ok := wechatbot[botKey]
		if !ok {
			continue
		}
		log.Info(fmt.Sprintf("start send msg to %s\n", botKey))
		textMsgOption := make([]text.TextMsgOption, 0)
		for _, user := range userList {
			textMsgOption = append(textMsgOption, text.MentionByUserid(user.WechatUUID))
		}
		err = robot.PushTextMessage(
			conf.AppConfig.MsgContent, textMsgOption...,
		)
		if err != nil {
			log.Error(fmt.Sprintf("send msg to bot(%s) error:%+v\n", botKey, err))
		}
	}

	return err
}
