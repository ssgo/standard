# 服务化架构标准化

### 日志标准化

##### 一般日志输出示例

```json
{
	"_logLevel": "info",
	"_logTime": 1552718260.2895439,
	"_logType": "Test",
	"level": 2
}
```

```json
{
	"_logLevel": "warning",
	"_logTime": 1552718260.28969,
	"_logType": "Test",
	"_traces": "/Volumes/Star/com.isstar/ssgo/standard/simple/Logger.go:40; /Volumes/Star/com.isstar/ssgo/standard/simple/Logger_test.go:20",
	"level": 3
}
```

```json
{
	"_logLevel": "error",
	"_logTime": 1552718260.289711,
	"_logType": "Test",
	"_traces": "/Volumes/Star/com.isstar/ssgo/standard/simple/Logger.go:44; /Volumes/Star/com.isstar/ssgo/standard/simple/Logger_test.go:21",
	"level": 4
}
```

##### 请求日志示例

```json
{
	"_logLevel": "info",
	"_logTime": 1552718260.2900941,
	"_logType": "request",
	"app": "appA",
	"authLevel": 1,
	"fromApp": "appB",
	"fromNode": "10.3.22.171:12334",
	"clientId": "HJDWAdaukhASd7",
	"clientIp": "59.32.113.241",
	"host": "abc.com",
	"method": "POST",
	"node": "10.3.22.178:32421",
	"path": "/users/{userId}/events",
	"priority": 2,
	"requestData": {
		"userId": 31123
	},
	"requestHeaders": {
		"Access-Token": "ab****fg"
	},
	"requestId": "udaHdhagy31Dd",
	"responseCode": 200,
	"responseData": {
		"events": null
	},
	"responseDataLength": 3401,
	"responseHeaders": {
		"XXX": "abc"
	},
	"sessionId": "8suAHDgsyakHU",
	"specialTag": true,
	"usedTime": 0.016
}
```
