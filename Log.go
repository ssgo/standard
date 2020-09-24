package standard

const LogTypeDebug = "debug"
const LogTypeInfo = "info"
const LogTypeWarning = "warning"
const LogTypeError = "error"
const LogTypeUndefined = "undefined"     // 无法识别的
const LogTypeDb = "db"                   // 数据库操作
const LogTypeDbError = "dbError"         // 数据库错误
const LogTypeServer = "server"           // 服务信息
const LogTypeServerError = "serverError" // 服务错误
const LogTypeTask = "task"               // 任务
const LogTypeMonitor = "monitor"         // 监控
const LogTypeStatistic = "statistic"     // 统计
const LogTypeRequest = "request"         // 服务请求

const LogDefaultSensitive = "phone,password,secure,token,accessToken" // 默认脱敏字段
const LogEnvLevel = "LOG_LEVEL"                                       // 日志级别，debug、info、warning、error，默认值：info
const LogEnvFile = "LOG_FILE"                                         // 日志文件，默认输出到 stdout
const LogEnvSensitive = "LOG_SENSITIVE"                               // 日志脱敏字段，默认值：standard.LogDefaultSensitive
const LogEnvRegexSensitive = "LOG_REGEXSENSITIVE"                     // 日志脱敏正则，无特殊情况不建议使用

type BaseLog struct {
	LogName    string // 日志名称
	LogType    string // 日志类型
	LogTime    string // 日志时间，格式为 RFC3339Nano "2006-01-02T15:04:05.999999999Z07:00"
	TraceId    string // 跟踪ID，同一过程中的日志具有相同的traceId
	ImageName  string // docker镜像名字
	ImageTag   string // docker镜像Tag
	ServerName string // 服务主机(宿主机器)名称
	ServerIp   string // 服务主机(宿主机器)Ip地址
	Extra      map[string]interface{}
}

type DebugLog struct {
	BaseLog
	Debug      string   // 调试信息
	CallStacks []string // 调用堆栈
}

type InfoLog struct {
	BaseLog
	Info string // 信息
}

type WarningLog struct {
	BaseLog
	Warning    string   // 警告信息
	CallStacks []string // 调用堆栈
}

type ErrorLog struct {
	BaseLog
	Error      string   // 错误信息
	CallStacks []string // 调用堆栈
}

type DBLog struct {
	BaseLog
	DbType    string  // 数据库类型，例如：mysql、oracle、redis...
	Dsn       string  // 连接信息
	Query     string  // 请求内容
	QueryArgs string  // 请求参数，会变化的部分应该记录在此
	UsedTime  float32 // 处理请求花费的时间，格式为float32，单位毫秒
}

type DBErrorLog struct {
	ErrorLog
	DBLog
}

type ServerLog struct {
	InfoLog
	App       string // 运行什么应用
	Weight    int    // 服务的权重
	Node      string // 运行在哪个节点（ip:port）
	Proto     string // 工作协议，例如：http1.1、http2.0、h2c
	StartTime string // 日志时间，格式为 RFC3339Nano "2006-01-02T15:04:05.999999999Z07:00"
}

type ServerErrorLog struct {
	ErrorLog
	App       string // 运行什么应用
	Weight    int    // 服务的权重
	Node      string // 运行在哪个节点（ip:port）
	Proto     string // 工作协议，例如：http1.1、http2.0、h2c
	StartTime string // 日志时间，格式为 RFC3339Nano "2006-01-02T15:04:05.999999999Z07:00"
}

type TaskLog struct {
	BaseLog
	Name      string                 // 任务名称
	Args      map[string]interface{} // 任务参数
	Succeed   bool                   // 是否成功
	Node      string                 // 运行在哪个节点（ip:port）
	StartTime string                 // 日志时间，格式为 RFC3339Nano "2006-01-02T15:04:05.999999999Z07:00"
	UsedTime  float32                // 处理任务花费的时间，格式为float32，单位毫秒
	Memo      string                 // 备注
}

type MonitorLog struct {
	BaseLog
	Name       string  // 监控项目
	Target     string  // 监控目标
	TargetInfo string  // 目标信息，例如：DNS、URL
	Expect     string  // 预期结果
	Result     string  // 实际结果
	Succeed    bool    // 是否成功
	UsedTime   float32 // 处理请求花费的时间，格式为float32，单位毫秒
	Memo       string  // 备注
}

type StatisticLog struct {
	BaseLog
	ServerId  string  // 服务编号（用于跟踪哪一个服务）
	App       string  // 运行什么应用
	Name      string  // 统计项目
	StartTime string  // 日志时间，格式为 RFC3339Nano "2006-01-02T15:04:05.999999999Z07:00"
	EndTime   string  // 日志时间，格式为 RFC3339Nano "2006-01-02T15:04:05.999999999Z07:00"
	Total     uint    // 总次数
	Failed    uint    // 失败次数
	AvgTime   float32 // 平均用时
	MinTime   float32 // 最少用时
	MaxTime   float32 // 最多用时
}

type RequestLog struct {
	BaseLog
	ServerId           string // 服务编号（用于跟踪哪一个服务）
	App                string // 应用名
	Node               string // 处理请求的节点，ip:port
	ClientIp           string // 真实的用户IP，通过 X-Real-IP 续传
	FromApp            string // 调用方应用
	FromNode           string // 调用方节点，格式 ip:port
	UserId             string // 客户唯一编号，通过 X-Client-ID 续传
	DeviceId           string
	ClientAppName      string
	ClientAppVersion   string
	SessionId          string                 // 会话唯一编号，通过 X-Session-ID 续传
	RequestId          string                 // 请求唯一编号，通过 X-Request-ID 续传
	Host               string                 // 真实用户请求的Host，通过 X-Host 续传
	Scheme             string                 // http scheme, http or https
	Proto              string                 // http proto, 1.1 or 2.0
	AuthLevel          int                    // 验证级别，用来校验用户是否有权限访问
	Priority           int                    // 优先级，用来在服务故障时进行自动降级处理
	Method             string                 // 请求的方法
	Path               string                 // 请求的路径，不包括GET参数部分，如果有PATH参数应该记录定义的PATH
	RequestHeaders     map[string]string      // 请求头，排除掉指定不需要信息后的所有头部内容，敏感数据应脱敏
	RequestData        map[string]interface{} // 请求的数据内容，JSON对象，集合类型仅记录少量内容，敏感数据应脱敏，非对象内容过大应做截取
	UsedTime           float32                // 处理请求花费的时间，格式为float32，单位毫秒
	ResponseCode       int                    // 应答代码，200 1000+ 正常应答，201～399，1～199  600～999 特殊应答，<1 异常应答
	ResponseHeaders    map[string]string      // 应答头，排除掉指定不需要信息后的所有头部内容，敏感数据应脱敏
	ResponseDataLength uint                   // 应答的数据长度
	ResponseData       string                 // 指定要记录的数据内容，JSON对象，集合类型仅记录少量内容，敏感数据应脱敏，非对象内容不进行记录
}
