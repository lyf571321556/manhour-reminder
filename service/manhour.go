package service

import (
	"encoding/json"
	"fmt"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"io/ioutil"
	"net/http"
)

var (
	login_auth = ""
)

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
				"quick":      "today",
			},
			"timeSeriesWithWorkDays": map[string]interface{}{
				"timeField":  "users.manhours.startTime",
				"valueField": "users.manhours.hours",
				"unit":       "day",
				"quick":      "today",
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

func FetchManhourByUUIDAndDepartmentUUID(url string, auth User, departmentUUID string, userUUIDS []string) (manhours []interface{}, err error) {
	manhourGql := generateGqlForManhour(userUUIDS, departmentUUID)
	var gqlJson []byte
	if gqlJson, err = json.Marshal(manhourGql); err != nil {
		return manhours, err
	}
	request, err := api.NewRequest(http.MethodPost, url, gqlJson)
	if err != nil {
		return manhours, err
	}

	request.Header.Add("Ones-User-Id", auth.UserId)
	request.Header.Add("Ones-Auth-Token", auth.Token)
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
			return fmt.Errorf("parse error: %w\nraw response: %s", err, rawResp)
		}

		buckets := dataMap["data"]["buckets"]
		if bucketsMapList, ok := buckets.([]interface{}); ok {
			for _, bucket := range bucketsMapList {
				println(bucket)
				str, err := json.Marshal(bucket)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("map to json", string(str))
			}
		}
		println(buckets)
		return nil
	})

	return manhours, err
}
