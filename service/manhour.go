package service

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
				"quick":      "last_7_days",
			},
			"timeSeriesWithWorkDays": map[string]interface{}{
				"timeField":  "users.manhours.startTime",
				"valueField": "users.manhours.hours",
				"unit":       "day",
				"quick":      "last_7_days",
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
