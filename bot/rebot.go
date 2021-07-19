package bot

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/config"
	"github.com/lyf571321556/manhour-reminder/log"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"github.com/lyf571321556/qiye-wechat-bot-api/bot"
	"github.com/lyf571321556/qiye-wechat-bot-api/text"
)

var wechatbot map[string]api.QiyeWechatBot

func InitBot() {
	wechatbot = make(map[string]api.QiyeWechatBot, 0)
	for _, botInfo := range config.AppConfig.RobotList {
		bot := bot.NewQiyeWechatBot(botInfo.RobotKey)
		api.SetDebug(true)
		wechatbot[botInfo.RobotKey] = bot
	}
}

func SendMsgToUser(auth service.AuthInfo) (err error) {
	list, err := service.FetchNeedToRemindUserlist(auth)
	if err != nil {
		log.Error(fmt.Sprintf("fetchNeedToRemindUserlist error:%+v\n", err))
		return err
	}

	for botKey, userList := range list {
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
			config.AppConfig.MsgContent, textMsgOption...,
		)
		if err != nil {
			log.Error(fmt.Sprintf("send msg to bot(%s) error:%+v\n", botKey, err))
		}
	}

	return err
}
