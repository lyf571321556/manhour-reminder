package service

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/conf"
	"github.com/lyf571321556/manhour-reminder/log"
)

func FetchNeedToRemindUserlist(auth AuthInfo) (list map[string][]UserInfo, err error) {
	list = make(map[string][]UserInfo, 0)
	for _, botInfo := range conf.AppConfig.RobotList {
		reminedUserInBot, err := fetchNeedRemidedUserInBot(auth, botInfo)
		if err != nil {
			log.Error(fmt.Sprintf("Bot(%s-%s) occur error:%+v\n", botInfo.RobotKey, botInfo.RobotName, err))
			continue
		}
		if len(reminedUserInBot) > 0 {
			list[botInfo.RobotKey] = reminedUserInBot
		}
	}
	return list, err
}

func fetchNeedRemidedUserInBot(auth AuthInfo, robotinfo conf.RobotInfo) (reminedUserInBot []UserInfo, err error) {
	reminedUserInBot = make([]UserInfo, 0)
	fetchManhourUrl := fmt.Sprintf("%s%s", conf.AppConfig.OnesProjectUrl, fmt.Sprintf(ITEMS_GQL, conf.AppConfig.TeamUUID))
	departmentUUID := robotinfo.DepartmentUUID
	userUUIDs := make([]string, 0)
	onesUserIdToWechatUserIdMapping := make(map[string]string, 0)
	for _, mapping := range robotinfo.UserMappings {
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
