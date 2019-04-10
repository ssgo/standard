package standard

const LogLevelDebug = "debug"
const LogLevelInfo = "info"
const LogLevelWarning = "warning"
const LogLevelError = "error"

const LogFieldType = "logType"                  // 日志类型
const LogFieldTime = "logTime"                  // 日志时间，格式为float64，单位秒
const LogFieldLevel = "logLevel"                // 日志级别
const LogFieldTraces = "traces"                 // 调用跟踪，以 "; " 间隔的字符串
const LogTypeUndefined = "undefined"            // 不符合规则的日志
