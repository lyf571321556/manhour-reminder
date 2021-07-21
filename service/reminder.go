package service

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/conf"
	"github.com/lyf571321556/manhour-reminder/log"
	"time"
)

func FetchNeedToRemindUserlist(auth AuthInfo) (list map[string][]*ManhourInfo, err error) {
	list = make(map[string][]*ManhourInfo, 0)
	for _, botInfo := range conf.AppConfig.RobotList {
		UserManhoursInBot, err := fetchUserManhoursInBot(auth, botInfo)
		if err != nil {
			log.Error(fmt.Sprintf("Bot(%s-%s) occur error:%+v\n", botInfo.RobotKey, botInfo.RobotName, err))
			continue
		}
		var manhours []*ManhourInfo = make([]*ManhourInfo, 0)
		for _, manhour := range UserManhoursInBot {
			manhours = append(manhours, manhour)
		}
		list[botInfo.RobotKey] = manhours
	}
	return list, err
}

func fetchUserManhoursInBot(auth AuthInfo, robotinfo conf.RobotInfo) (userUUIDToManhoutMap map[string]*ManhourInfo, err error) {
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
		return manhourMapping, err
	}
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	for _, manhourinfo := range manhourMapping {
		times := manhourinfo.ActualHoursSeries.Times
		values := manhourinfo.ActualHoursSeries.Values
		for index, time := range times {
			if time == currentDate && values[index] == 0 {
				manhourinfo.User.WechatUUID = onesUserIdToWechatUserIdMapping[manhourinfo.User.UUID]
				manhourinfo.User.Reminded = true
			}
		}
	}
	return manhourMapping, err
}
