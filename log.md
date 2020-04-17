# 日志格式说明

## 1.日志概述
日志均由日志公用部分和日志内容体两部分组成。
日志头是日志格式中通用信息，每种类型日志可能存在差异。日志内容体格式由日志类型(logType)决定，不同日志类型可能内容格式不同。

## 2. 应用服务日志
应用服务日志，用于调试或者跟踪以及出错分析，主要使用者为开发、测试、运维等人员。

### 2.1 应用服务日志日志公用部分

注：以下是日志格式中,json关键字说明
日志头部分

| 名称 | 字段类型| 描述|  必填项 or 可选项 | 备注 |
| :------ | :------ |:----------- | :------------ | :------|
| logName |string |日志名称| 可选项 | 日志所属的服务名或者docker的镜像名|
| logType |string |日志类型| 必填项 | 按照功能区分, 如: request, server, serverError, db, dbError, statistics, system(默认), monitor, task|
| logTime |string | 日志时间 | 必填项 | 格式为: RFC3339Nano "2006-01-02T15:04:05.999999999Z07:00"|
| logLevel |string | 日志级别 | 必填项 | 内容必须为以下：debug, info, warning, error |
| traceId |string |日志Id, 如有请填写。| 可选项 | <br>最好使用gateway统一生成的UUId类标识,算法可以是[snowflake](https://blog.twitter.com/2010/announcing-snowflake) 和 uuId[rfc4122](https://tools.ietf.org/html/rfc4122) |
| imageName |string | docker镜像名字 | 可选项|  收集程序补全|
| imageTag | string | docker镜像Tag | 可选项 | 收集程序补全|
| serverName |string | 服务主机(宿主机器)名称 | 可选项 |  收集程序补全|
| serverIp | string | 服务主机(宿主机器)Ip地址 | 可选项 | 收集程序补全|

收集程序负责补全信息。

### 2.2 应用服务日志格式说明

一条应用服务日志为例输出结构如下：

    2019/06/13 09:38:36.572087 {"app":"","authLevel":0,"clientId":"","clientIp":"127.0.0.1","fromApp":"","fromNode":"","host":"127.0.0.1:8081","logTime":1560389916.572087,"logType":"request","method":"POST","node":"10.59.10.143:8081","p
    ath":"/sskey","priority":0,"proto":"1.1","requestData":{"dawei":"JRygtTB3qLRW6XDwUXCU4JjiReKqoXwvzLMgyJwkDoAPI98VCE5Bm8k+l3MfefoTcqxfRYNq+D4yx1I65B+DkdDye3p0ue/lhRQNGO/jQT1CJaPhk1MhYCu6jOU0ExSw","ldw":"ulbyB3aLloskgqNlrQFLWmq8CcRvTW
    oqvIKY7AOF7wwHnrfHjjvkdWHodfwmflEyvIRKzj06Bg1Cy+8xC1+akN86tw+3xGWrGhIe83qKHQNttULmC3bCepjy6L4rsePS"},"requestHeaders":{"Content-Length":"277","Content-Type":"application/json","User-Agent":"Go-http-client/1.1","X-Host":"127.0.0.1:80
    81","X-Real-Ip":"127.0.0.1","X-Request-Id":"d0SGlgd40uuAUZCDfiLP85","X-Scheme":"http"},"requestId":"d0SGlgd40uuAUZCDfiLP85","responseCode":200,"responseData":true,"responseDataLength":4,"responseHeaders":{},"scheme":"http","serverId
    ":"T1uBk71ebv","sessionId":"","traceId":"d0SGlgd40uuAUZCDfiLP85","usedTime":0,"type":"ACCESS"}

包含两个部分：

* 前置时间

格式为：Y/m/d H:i:s.us 年/月/日 24时/分/秒.微秒。

* 日志主体

格式为：日志明细内容的json格式输出。

总体格式为："前置时间 json日志主体"。

以下类型中对于主要字段，推荐但不限制表格中字段。个性化字段信息请自行添加，满足需求即可。

#### 2.2.1 request类型
一般用于记录服务处理请求、响应日志。

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| sessionId | string | session标识,如有请填写 | 可选项 |
| clientId | string | 客户Id,如有请填写 | 可选项 |
| clientIp | string | 客户端Ip地址 | 必填项 |
| app | string| 应用名称, 如有请填写 | 可选项 |
| host | string |请求主机 | 必填项 |
| method | string | 请求方法,如"GET","POST","PUT"等| 必填项 |
| path | string | 请求路径 | 必填项 |
| handle | string | 处理请求方法 | 必填项 |
| requestId | string | 请求Id | 可选项 |
| requestHeaders | string | 请求数据headers| 可选项 |
| requestData | string| 请求数据（可能需要脱敏） | 可选项 |
| priority | number | 业务优先级,主要用于降级场景,默认为0-99,越高越大 | 可选项 |
| authLevel | number | 权限级别,主要不同业务的安全控制,默认为0-9,越高权限越大 | 可选项 |
| responseCode | string | 返回code | 必填项 |
| responseData | string | 返回数据,如有请填写 | 可选项 |
| responseDataLength | string |返回数据长度,如有请填写 | 可选项 |
| responseHeaders | string | 返回数据headers| 可选项 |
| usedTime | number | 处理请求花费时间(单位秒精确到毫秒) | 可选项 |
| fromApp | string| 来源应用,如有请填写 | 可选项 |
| fromNode | string | 来源节点,如有请填写 | 可选项 |

* 例子1：正常处理

```json
{
   "logLevel": "info",
	"logTime": 1552718260.2895439,
	"traceId": "1119119361907363840",
	"requestId": "1119119361907363840",
	"logType": "request",
	"app": "appA",
	"authLevel": 1,
	"fromApp": "appB",
	"fromNode": "10.3.22.171:12334",
	"service":"user",
	"clientId": "HJDWAdaukhASd7",
	"clientIp": "59.32.113.241",
	"host": "abc.com",
	"method": "POST",
	"node": "10.3.22.178:32421",
	"path": "/users/{userId}/events",
	"handle":"handleEvents",
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
	"usedTime": 1.160
}
```
* 例子2：处理异常

```json
{
   "logLevel": "error",
	"logTime": 1552718260.2895439,
	"logType": "request",
	"traceId": "1119119361907363840",
	"requestId": "1119119361907363840",
	"app": "appA",
	"authLevel": 1,
	"fromApp": "appB",
	"fromNode": "10.3.22.171:12334",
	"clientId": "HJDWAdaukhASd7",
	"clientIp": "59.32.113.241",
	"module":"user",
	"host": "abc.com",
	"method": "GET",
	"node": "10.3.22.178:32421",
	"path": "/users/event/{event_Id}",
	"priority": 3,
	"requestHeaders": {
		"Access-Token": "ab****fg"
	},
	"requestId": "udaHdhagy31Dd",
	"msg": "failed to connect MySQL",
	"responseCode": 500,
	"responseHeaders": {
		"XXX": "abc"
	},
	"sessionId": "8suAHDgsyakHU",
	"specialTag": true,
	"usedTime": 306
}
```


#### 2.2.2 server类型
server类型用于记录服务运行情况

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| weight | number |权重 | 可选项 |
| node | string | 运行节点（ip:port）,如有多个使用';'分隔 | 必填项 |
| proto  | string | 工作协议，例如：http1.1、http2.0、h2c | 必填项 |
| startTime | number | 服务启动时间 | 必填项 |
| msg | string | 服务启动信息 | 可选项 |

* 例子

```json
{
	"logLevel": "info",
	"logTime": 1552718260.2895439,
	"logType": "server",
	"weight": 80,
	"node": "172.16.1.1:8080",
	"proto": "h2c",
	"startTime": 1552718260.2910283,
	"msg":"learn is starting"
}
```

#### 2.2.3 serverError类型
server类型用于描述服务运行错误的情况信息

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| weight | number |权重 | 必填项 |
| node | string | 运行节点（ip:port）,如有多个使用';'分隔 | 可选项 |
| proto  | string | 工作协议，例如：http1.1、http2.0、h2c | 必填项 |
| startTime | number | 服务启动时间 | 可选项 |
| msg | string | 服务停止的错误信息 | 必填项 |

* 例子

```json
{
	"logLevel": "error",
	"logTime": 1552718268.2334243,
	"logType": "serverError",
	"weight": 80,
	"node": "172.16.1.1:8080",
	"proto": "h2c",
	"startTime": 1552718260.2910283,
	"msg": "connect to mysql timeout, learn is stopping"
}
```


#### 2.2.4 db类型
数据库操作日志，用于跟踪调试sql和nosql成功执行情况。

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| dbType | string | 数据库类型: mysql, redis, memcached, kafka | 必填项 |
| dsn | string | DB地址或者nosql地址如: "mysql://192.168.2.11:3306/test"; "redis://192.168.1.1:6379/0"; "mc://192.168.1.1:11211"; "kafka://192.168.1.1:9092,192.168.1.2:9092,192.168.1.3:9092/topic" | 必填项 |
| query | string | 执行sql语句或执行nosql命令 | 可选项 |
| args  | array | 执行sql语句或nosql命令的参数 | 可选项 |
| usedTime | number | 执行sql语句花费时间(单位秒精确到毫秒) | 必填项 |

* 例子1: 

```json
{
	"logLevel": "info",
	"logTime": 1552718260.2895439,
	"logType": "db",
	"dbType": "mysql",
	"dsn": "mysql://192.168.2.11:3306/test",
	"query": "select aa, bb from table where Id = ?",
	"args": [900],
	"usedTime": 2.500
}
```

* 例子2: 

```json
{
	"logLevel": "info",
	"logTime": 1552718260.2895439,
	"logType": "db",
	"dbType": "redis",
	"dsn": "redis://192.168.2.11:6379/0",
	"query": "HSET website mail",
	"args": ["mail.hfjy.com"],
	"usedTime": 2.500
}
```

* 例子3: 

```json
{
	"logLevel": "info",
	"logTime": 1552718260.2895439,
	"logType": "db",
	"dbType": "kafka",
	"dsn": "kafka://192.168.1.1:9092,192.168.1.2:9092,192.168.1.3:9092/topic",
	"query": "producer",
	"args": ["hfjy message"],
	"usedTime":2.500
}
```

#### 2.2.5 dbError类型
数据库操作日志，用于记录sql和nosql执行失败的情况。

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| dbType | string | 数据库类型: mysql, redis, memcached, kafka | 必填项 |
| dsn | string | DB地址或者nosql地址如: "mysql://192.168.2.11:3306/test"; "redis://192.168.1.1:6379/0"; "mc://192.168.1.1:11211"; "kafka://192.168.1.1:9092,192.168.1.2:9092,192.168.1.3:9092" | 必填项 |
| query | string | 执行sql语句或执行nosql命令 | 必填项 |
| args | array | 执行sql语句或nosql命令的参数 | 可选项 |
| msg | string | 执行sql语句出错时填写错误信息,如有错误码, 请按照"errorCode:msg"记录日志内容 | 必填项 |
| usedTime | number | 执行sql语句花费时间(单位秒精确到毫秒) | 可选项 |

* 例子

```json
{
	"logLevel": "error",
	"logTime": 1552718260.2895439,
	"logType": "dbError",
	"dbType": "mysql",
	"dsn": "mysql://192.168.2.11:3306/test",
	"query": "select aa, bb from notable where Id = ?",
	"args": [900],
	"msg": "Table doesn't exist",
	"usedTime": 2.500
}
```

#### 2.2.6 statistics统计资源类型
用于记录应用服务资源统计信息,比如: 服务的负载情况。 对于有特殊要求的场景使用。

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| serverId | string | 服务编号（用于跟踪哪一个服务）| 必填项 |
| name | string | 统计项目名字| 必填项 |
| startTime | number | 开始时间 | 必填项 |
| endTime | number | 结束时间 | 必填项 |
| total | number | 总次数| 可选项 |
| failed | number | 失败次数 | 可选项 |
| avgTime | number | 平均用时(单位秒精确到毫秒) | 可选项 |
| minTime | number | 最小用时(单位秒精确到毫秒) | 可选项 |
| maxTime | number | 最大用时(单位秒精确到毫秒) | 可选项 |

* 例子

```json
{
	"logTime": 1552718260.2895439,
	"logType": "statistics",
	"serverId": "learn",
	"name": "status",
	"startTime": 1552718200.2895439,
	"endTime": "1552718260.2895439",
	"total": 11216,
	"failed": 0,
	"avgTime": 1.050,
	"minTime": 1.030,
	"maxTime": 1.082,
}
```

#### 2.2.7 system类型
用于记录应用Service系统框架层日志信息,比如: 日志框架打开日志文件异常。

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| msg | string |日志信息,如有错误码, 请按照"errorCode:msg"记录日志内容| 必填项 |
| callStacks | array string |用于调用栈跟踪,如有请填写 |可选项 |

* 例子1:

```json
{
	"logLevel": "error",
	"logTime": "1552718260.2895439",
	"logType": "system",
	"msg":"1002:file does not exist"
}
```

* 例子2:

```json
{
    "logLevel":"debug",
    "logTime":"1552718260.2895439",
    "logType":"system",
    "callStacks":[
        "/Volumes/Star/com.isstar/ssgo/standard/simple/Logger_test.go:20"
    ],
    "msg":"debug information"
}
```

#### 2.2.8 monitor监控类型
用于监控项目情况。

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| name | string | 监控项目名字| 必填项 |
| target | string | 监控目标| 必填项 |
| targetInfo | string | 目标信息如：dsn, url| 必填项 |
| expect | string | 预期结果| 必填项 |
| result | string | 实际结果| 必填项 |
| isSuccess | bool | 是否成功| 必填项 |
| usedTime | number | 监控项目花费的时间(单位秒精确到毫秒)| 必填项 |
| memo | string | 备注 | 可选项 |


* 例子

```json
{
	"logTime": 1552718260.2895439,
	"logType": "monitor",
	"name": "service",
	"target": "learn",
	"targetInfo": "http://192.168.1.1:8000/ping",
	"expect": "pong",
	"Result": "pong",
	"isSuccess": true,
	"usedTime": 0.050,
}
```


#### 2.2.9 task类型
用于记录task执行情况。主要用于一个请求或处理涉及几个子任务的场景。

* 主要字段说明

|字段名称 |字段类型 |备注  |必填项 or 可选项  |
| :--- | :--- | :--- | :--- |
| name | string | task名字| 必填项 |
| args | string | 任务参数| 必填项 |
| success | bool | 是否成功| 必填项 |
| node | string | 运行在哪个节点（ip:port）| 必填项 |
| startTime | float64 | 任务开始时间，格式为float64，单位秒| 必填项 |
| usedTime | number | 处理task花费的时间(单位秒精确到毫秒)| 必填项 |
| memo | string | 备注 | 可选项 |

* 例子

```json
{
	"logTime": 1552718260.2895439,
	"logType": "task",
	"serverId": "learn",
	"name": "step1-get-users",
	"isSuccess": true,
	"usedTime": 0.550,
}
```



