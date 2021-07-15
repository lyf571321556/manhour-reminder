package main

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/config"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/lyf571321556/qiye-wechat-bot-api/text"
	"log"
)

func main() {
	msg := text.TextMsg{MsgType: "", ContentData: &text.ContentData{
		Content:             "",
		MentionedList:       nil,
		MentionedMobileList: nil,
	}}
	print(msg.MsgType)
	println(config.AppConfig.OnesProjectUrl)
	for _, botInfo := range config.AppConfig.BotList {
		for _, mapping := range botInfo.UserMappings {
			log.Printf("botName:%s,departmentuuid:%s,ones_userid:%s,wechat_userid:%s", botInfo.BotName, botInfo.DepartmentUUID, mapping.OnesUserid, mapping.WechatUserid)
		}
	}

	//获取token
	loginUrl := fmt.Sprintf("%s%s", config.AppConfig.OnesProjectUrl, service.AUTH_LOGIN)
	token, err := service.Login(loginUrl, "wuxingjuan@ones.ai", "juan1997")
	if err != nil {
		log.Fatal(err)
	}
	println(token)

}
