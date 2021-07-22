package robot

import (
	"bytes"
	"fmt"
	"github.com/lyf571321556/manhour-reminder/conf"
	"github.com/lyf571321556/manhour-reminder/log"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"github.com/lyf571321556/qiye-wechat-bot-api/bot"
	"github.com/lyf571321556/qiye-wechat-bot-api/text"
	"strconv"
)

var Wechatbot map[string]*bot.QiyeWechatBot

func InitBot(configPath string) (err error) {
	err = conf.Init(configPath)
	if err != nil {
		return err
	}
	err = log.InitLog()
	if err != nil {
		return err
	}
	Wechatbot = make(map[string]*bot.QiyeWechatBot, 0)
	for _, botInfo := range conf.AppConfig.RobotList {
		bot := bot.NewQiyeWechatBot(botInfo.RobotName, botInfo.RobotKey)
		api.SetDebug(true)
		Wechatbot[botInfo.RobotKey] = bot
	}
	return err
}

func StartCheckUsersManhourInEveryRobot(auth service.AuthInfo) (err error) {
	list, err := service.FetchNeedToRemindUserlist(auth)
	if err != nil {
		log.Error(fmt.Sprintf("fetchNeedToRemindUserlist error:%+v\n", err))
	}
	if len(list) > 0 {
		err = sendMsgToUser(list)
	}
	return err
}

func sendMsgToUser(uerMahoursInRobot map[string][]*service.ManhourInfo) (err error) {
	if uerMahoursInRobot == nil {
		return err
	}
	for botKey, manhours := range uerMahoursInRobot {
		robot, ok := Wechatbot[botKey]
		if !ok {
			continue
		}
		pushMsg, textMsgOption := buildPushMsgAndMentions(manhours)
		log.Info(fmt.Sprintf("start send msg to robot %s(%s)\n", robot.Name, botKey))
		err = robot.PushTextMessage(
			pushMsg, textMsgOption...,
		)
		if err != nil {
			log.Error(fmt.Sprintf("send msg to robot %s(%s) error:%+v\n", robot.Name, botKey, err))
		}
		log.Info(fmt.Sprintf("send msg to robot %s(%s) success\n", robot.Name, botKey))
	}
	return err
}

func buildPushMsgAndMentions(uerMahoursInRobot []*service.ManhourInfo) (pushMsg string, textMsgOption []text.TextMsgOption) {
	var msgManhoursTableContent bytes.Buffer
	msgManhoursTableContent.WriteString("工时提醒：请同学们根据工时记录检查是否及时登记工时~\n")
	remindedUsers := make([]service.UserInfo, 0)
	for index, manhour := range uerMahoursInRobot {
		if index == 0 {
			for timeindex, time := range manhour.ActualHoursSeries.Times {
				if timeindex > 0 {
					msgManhoursTableContent.WriteString(" 🕙 ")
				}
				msgManhoursTableContent.WriteString(time)
			}
			msgManhoursTableContent.WriteString("\n")
		}
		msgManhoursTableContent.WriteString(manhour.User.UserName)
		msgManhoursTableContent.WriteString("\n")
		for valueIndex, value := range manhour.ActualHoursSeries.Values {
			if valueIndex > 0 {
				msgManhoursTableContent.WriteString(" / ")
			}
			msgManhoursTableContent.WriteString(strconv.FormatFloat(value/100000, 'f', 1, 64))
		}
		msgManhoursTableContent.WriteString("\n")
		if manhour.User.Reminded {
			remindedUsers = append(remindedUsers, manhour.User)
		}
	}
	textMsgOption = make([]text.TextMsgOption, 0)
	if len(remindedUsers) > 0 {
		msgManhoursTableContent.WriteString("以下人员今天还未登记工时，请及时补充工时信息。")
	}
	for _, user := range remindedUsers {
		textMsgOption = append(textMsgOption, text.MentionByUserid(user.WechatUUID))
	}
	return msgManhoursTableContent.String(), textMsgOption

}
