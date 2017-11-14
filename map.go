package main

// InitMapper create uri mapper
func InitMapper() map[string]string {
	var uriMapper = map[string]string{
		"/api/device-mgt/device/list":                       "mock/device-mgt/device/list.json",
		"/api/device-mgt/device/statics":                    "mock/device-mgt/device/statics.json",
		"/api/device-mgt/device/inventory":                  "mock/device-mgt/device/inventory.json",
		"/api/device-mgt/device/sensor/list":                "mock/device-mgt/device/sensor/list.json",
		"/api/device-mgt/device/sensor/statics":             "mock/device-mgt/device/sensor/statics.json",
		"/api/device-mgt/device/sensor":                     "mock/device-mgt/device/sensor/sensor.json",
		"/api/device-mgt/device/sensor/history/statics":     "mock/device-mgt/device/sensor/historyStatics.json",
		"/api/device-mgt/device/sensor/history":             "mock/device-mgt/device/sensor/history.json",
		"/api/device-mgt/device/sel/list":                   "mock/device-mgt/device/sel/list.json",
		"/api/device-mgt/device/sel/statics":                "mock/device-mgt/device/sel/statics.json",
		"/api/device-mgt/device/config/device/list":         "mock/device-mgt/device/config/deviceList.json",
		"/api/device-mgt/device/config/service":             "mock/device-mgt/device/config/service.json",
		"/api/device-mgt/device/config/service/suggestions": "mock/device-mgt/device/config/suggestion.json",
		"/api/device-mgt/device/config/tag/list":            "mock/device-mgt/device/config/tagList.json",
		"/api/device-mgt/monitor/template/list":             "mock/device-mgt/device/monitor/templateList.json",
		"/api/device-mgt/firmware/task/list":                "mock/firmware/taskList.json",
		"/api/device-mgt/firmware/image/list":               "mock/firmware/imageList.json",
		"/api/device-mgt/firmware/template/list":            "mock/firmware/templateList.json",
		"/api/device-mgt/firmware/list":                     "mock/firmware/firmwareList.json",
		"/api/user/list":                                    "mock/user/userList.json",
		"/api/settings/system":                              "mock/setting/system.json",
		"/api/system/alert/list":                            "mock/system/alert.json",
		"/api/system/user-log/list":                         "mock/system/userLog.json",
		"/api/system/sys-log/list":                          "mock/system/sysLog.json",
		"/api/system/lang":                                  "mock/system/lang.json",
		// "/api/system/job/list":                              "mock/system/job.json",
	}
	return uriMapper
}
