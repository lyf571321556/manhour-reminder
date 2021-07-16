package service

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/config"
	"log"
)

func FetchNeedToRemindUserlist(auth AuthInfo) (list map[string][]UserInfo, err error) {
	list = make(map[string][]UserInfo, 0)
	for _, botInfo := range config.AppConfig.BotList {
		reminedUserInBot, err := fetchNeedRemidedUserInBot(auth, botInfo)
		if err != nil {
			log.Printf("Bot(%s-%s) occur error:%w\n", botInfo.BotKey, botInfo.BotName, err)
			continue
		}
		if len(reminedUserInBot) > 0 {
			list[botInfo.BotKey] = reminedUserInBot
		}
	}
	return list, err
}

func fetchNeedRemidedUserInBot(auth AuthInfo, botinfo config.BotInfo) (reminedUserInBot []UserInfo, err error) {
	reminedUserInBot = make([]UserInfo, 0)
	fetchManhourUrl := fmt.Sprintf("%s%s", config.AppConfig.OnesProjectUrl, fmt.Sprintf(ITEMS_GQL, config.AppConfig.TeamUUID))
	departmentUUID := botinfo.DepartmentUUID
	userUUIDs := make([]string, 0)
	onesUserIdToWechatUserIdMapping := make(map[string]string, 0)
	for _, mapping := range config.AppConfig.BotList[0].UserMappings {
		userUUIDs = append(userUUIDs, mapping.OnesUserid)
		onesUserIdToWechatUserIdMapping[mapping.OnesUserid] = mapping.WechatUserid
	}

	manhourMapping, err := FetchManhourByUUIDAndDepartmentUUID(fetchManhourUrl, auth, departmentUUID, userUUIDs)
	if err != nil {
		return reminedUserInBot, err
	}

	for _, manhourinfo := range manhourMapping {
		if manhourinfo.getActualHours() < 8 {
			manhourinfo.User.WechatUUID = onesUserIdToWechatUserIdMapping[manhourinfo.User.UUID]
			reminedUserInBot = append(reminedUserInBot, manhourinfo.User)
		}
	}
	return reminedUserInBot, err
}
