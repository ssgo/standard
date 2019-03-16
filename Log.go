package standard

import (
	"time"
)

type LogLevelType int

const LogDebug LogLevelType = 1
const LogInfo LogLevelType = 2
const LogWarning LogLevelType = 3
const LogError LogLevelType = 4

const LogDebugName = "debug"
const LogInfoName = "info"
const LogWarningName = "warning"
const LogErrorName = "error"

const LogTypeFieldName = "_logType"             // 日志类型
const LogTimeFieldName = "_logTime"             // 日志时间，格式为float64，单位秒
const LogLevelFieldName = "_logLevel"           // 日志级别
const LogTracesFieldName = "_traces"            // 调用跟踪，以 "; " 间隔的字符串
const LogEncodingErrorType = "logEncodingError" // 数据无法序列化为JSON时的日志类型
const LogRequestType = "request"                // 服务请求日志类型

const LogRequestAppFieldName = "app"                         // 应用名
const LogRequestNodeFieldName = "node"                       // 处理请求的节点，ip:port
const LogRequestClientIpFieldName = "clientIp"               // 真实的用户IP，通过 X-Real-IP 续传
const LogRequestCallerFieldName = "caller"                   // 调用方，格式 app:ip:port
const LogRequestClientIdFieldName = "clientId"               // 客户唯一编号，通过 X-Client-ID 续传
const LogRequestSessionIdFieldName = "sessionId"             // 会话唯一编号，通过 X-Session-ID 续传
const LogRequestRequestIdFieldName = "requestId"             // 请求唯一编号，通过 X-Request-ID 续传
const LogRequestHostFieldName = "host"                       // 真实用户请求的Host，通过 Host 续传
const LogRequestAuthLevelFieldName = "authLevel"             // 验证级别，用来校验用户是否有权限访问
const LogRequestPriorityFieldName = "priority"               // 优先级，用来在服务故障时进行自动降级处理
const LogRequestMethodFieldName = "method"                   // 请求的方法
const LogRequestPathFieldName = "path"                       // 请求的路径，不包括GET参数部分，如果有PATH参数应该记录定义的PATH
const LogRequestRequestHeadersFieldName = "requestHeaders"   // 请求头，排除掉指定不需要信息后的所有头部内容，敏感数据应脱敏
const LogRequestArgsFieldName = "requestData"                // 请求的数据内容，JSON对象，集合类型仅记录少量内容，敏感数据应脱敏，非对象内容过大应做截取
const LogRequestUsedTimeFieldName = "usedTime"               // 处理请求花费的时间，格式为float32，单位毫秒
const LogRequestStatusFieldName = "responseCode"             // 应答代码，200 1000+ 正常应答，201～399，1～199  600～999 特殊应答，<1 异常应答
const LogRequestResponseHeadersFieldName = "responseHeaders" // 应答头，排除掉指定不需要信息后的所有头部内容，敏感数据应脱敏
const LogRequestOutLenFieldName = "responseDataLength"       // 应答的数据长度
const LogRequestResultFieldName = "responseData"             // 指定要记录的数据内容，JSON对象，集合类型仅记录少量内容，敏感数据应脱敏，非对象内容不进行记录

func MakeLogTime(time time.Time) float64 {
	return float64(time.UnixNano()) / 1e9
}

func MakeUesdTime(startTime, endTime time.Time) float32 {
	return float32(endTime.UnixNano()-startTime.UnixNano()) / 1e6
}
