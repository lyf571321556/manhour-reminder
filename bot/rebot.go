package bot

import (
	"github.com/lyf571321556/manhour-reminder/config"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"github.com/lyf571321556/qiye-wechat-bot-api/bot"
	"github.com/lyf571321556/qiye-wechat-bot-api/text"
	"log"
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
		log.Printf("fetchNeedToRemindUserlist error:%w\n", err)
		return err
	}

	for botKey, userList := range list {
		robot, ok := wechatbot[botKey]
		if !ok {
			continue
		}
		log.Printf("start send msg to %s\n", botKey)
		textMsgOption := make([]text.TextMsgOption, 0)
		//var content bytes.Buffer
		for _, user := range userList {
			textMsgOption = append(textMsgOption, text.MentionByUserid(user.WechatUUID))
			//content.WriteString(fmt.Sprintf("%s,记得登记工时.\n", user.UserName))
		}
		err = robot.PushTextMessage(
			config.AppConfig.MsgContent, textMsgOption...,
		)
		if err != nil {
			log.Printf("send msg to bot(%s) error:%w\n", botKey, err)
		}
	}

	return err
}
