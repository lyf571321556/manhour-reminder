package service

import (
	"encoding/json"
	"fmt"
	"github.com/lyf571321556/manhour-reminder/log"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"io/ioutil"
	"net/http"
)

type ManhourInfo struct {
	ActualHours       int64                 `json:"actualHours"`
	User              UserInfo              `json:"columnField"`
	ActualHoursSeries ActualHoursSeriesInfo `json:"actualHoursSeries"`
}

type ActualHoursSeriesInfo struct {
	Times  []string  `json:"times"`
	Values []float64 `json:"values"`
}

type UserInfo struct {
	UUID       string `json:"uuid"`
	UserName   string `json:"name"`
	WechatUUID string `json:"wechat_uuid"`
	Reminded   bool
}

func (manhourinfo *ManhourInfo) getActualHours() int64 {
	return manhourinfo.ActualHours / 100000
}

func (manhourinfo *ManhourInfo) generateRemindDesc() string {
	if manhourinfo.getActualHours() >= 0 && manhourinfo.getActualHours() < 8 {
		return fmt.Sprintf("%s,记得登记工时。\n", manhourinfo.User.UserName)
	}
	return ""
}

func generateGqlForManhour(userids []string, departmentUUID string) interface{} {
	qglQuery := map[string]interface{}{
		"query": "query QUERY_USERS($groupBy: GroupBy, $orderBy: OrderBy, $timeSeries: TimeSeriesArgs, $timeSeriesWithWorkDays: TimeSeriesArgs, $actualHoursSum: String, $filter: Filter, $columnSource: Source) {\n  buckets(groupBy: $groupBy, orderBy: $orderBy, filter: $filter) {\n    ...ResourceLoadBucketFragment\n  }\n}\n\nfragment UserSimple on User {\n  key\n  uuid\n  name\n  avatar\n  email\n}\n\nfragment ResourceLoadBucketFragment on Bucket {\n  key\n  columnField: aggregateUser(source: $columnSource) {\n    ...UserSimple\n  }\n  actualHours: baseActualHoursTotal\n  estimatedHours: baseEstimatedHoursTotal\n  estimatedHoursSeries(timeSeries: $timeSeries) {\n    times\n    values\n  }\n  actualHoursSeries(timeSeries: $timeSeries) {\n    times\n    values\n  }\n  remainingWorkingHoursSeries: baseActualHoursRemainingWorkingHoursSeries {\n    times\n    values\n  }\n  standardWorkingHoursSeries(timeSeries: $timeSeriesWithWorkDays) {\n    times\n    values\n  }\n  workloadRateSeries: baseActualHoursWorkloadRateSeries(timeSeries: $timeSeries) {\n    times\n    values\n  }\n  averageWorkloadRate: baseActualHoursAverageWorkloadRate\n}\n",
		"variables": map[string]interface{}{
			"groupBy": map[string]interface{}{
				"users": map[string]interface{}{
					"uuid": map[string]interface{}{},
				},
			},
			"orderBy": map[string]interface{}{
				"aggregateUser": map[string]interface{}{
					"namePinyin": "ASC",
				},
			},
			"filter": map[string]interface{}{
				"users": map[string]interface{}{
					"status_in": []string{
						"normal",
					},
					"departments_in": []string{
						departmentUUID,
					},
					"uuid_in": userids,
				},
			},
			"timeSeries": map[string]interface{}{
				"timeField":  "users.manhours.startTime",
				"valueField": "users.manhours.hours",
				"unit":       "day",
				"quick":      "this_week",
			},
			"timeSeriesWithWorkDays": map[string]interface{}{
				"timeField":  "users.manhours.startTime",
				"valueField": "users.manhours.hours",
				"unit":       "day",
				"quick":      "this_week",
				"constant":   800000,
				"workdays": []string{
					"Mon",
					"Tue",
					"Wed",
					"Thu",
					"Fri",
				},
			},
			"columnSource": "uuid",
			"refreshStamp": 0,
		},
	}
	return qglQuery
}

func FetchManhourByUUIDAndDepartmentUUID(url string, auth AuthInfo, departmentUUID string, userUUIDS []string) (userUUIDToManhoutMap map[string]*ManhourInfo, err error) {
	manhourGql := generateGqlForManhour(userUUIDS, departmentUUID)
	var gqlJson []byte
	if gqlJson, err = json.Marshal(manhourGql); err != nil {
		return userUUIDToManhoutMap, err
	}
	request, err := api.NewRequest(http.MethodPost, url, gqlJson)
	if err != nil {
		return userUUIDToManhoutMap, err
	}

	request.Header.Add("Ones-User-Id", auth.UserId)
	request.Header.Add("Ones-Auth-Token", auth.Token)
	userUUIDToManhoutMap = make(map[string]*ManhourInfo, 0)
	_, err = api.ExecuteHTTP(request, func(resp *http.Response) error {
		rawResp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error:%s", string(rawResp))
		}

		dataMap := make(map[string]map[string]interface{})
		if err = json.Unmarshal(rawResp, &dataMap); err != nil {
			return fmt.Errorf("parse error: %+v\nraw response: %s", err, rawResp)
		}
		buckets := dataMap["data"]["buckets"]
		if bucketsMapList, ok := buckets.([]interface{}); ok {
			for _, bucket := range bucketsMapList {
				var manhourInfo = new(ManhourInfo)
				bytes, err := json.Marshal(bucket)
				if err != nil {
					log.Error(err.Error())
					continue
				}
				err = json.Unmarshal(bytes, manhourInfo)
				if err != nil {
					log.Error(err.Error())
					log.Error(err.Error())
					continue
				}
				userUUIDToManhoutMap[manhourInfo.User.UUID] = manhourInfo
			}
		}
		return err
	})

	return userUUIDToManhoutMap, err
}
