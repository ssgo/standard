package standard

const LogTypeDebug = "debug"
const LogTypeInfo = "info"
const LogTypeWarning = "warning"
const LogTypeError = "error"
const LogTypeUndefined = "undefined" // 无法识别的日志类型

const LogTypeDb = "db"           // 数据库操作日志类型
const LogTypeDbError = "dbError" // 数据库错误日志类型

const LogTypeStatistic = "statistic" // 统计日志类型

const LogTypeRequest = "request" // 服务请求日志类型

type BaseLog struct {
	LogType string  // 日志类型
	LogTime float64 // 日志时间，格式为float64，单位秒
	TraceId string  // 跟踪ID，同一过程中的日志具有相同的traceId
	Extra   map[string]interface{}
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
	DbType   string        // 数据库类型，例如：mysql、oracle、redis...
	Dsn      string        // 连接信息
	Query    string        // 请求内容
	Args     []interface{} // 请求参数，会变化的部分应该记录在此
	UsedTime float32       // 处理请求花费的时间，格式为float32，单位毫秒
}

type DBErrorLog struct {
	ErrorLog
	DBLog
}

type StatisticLog struct {
	BaseLog
	Project   string  // 统计项目
	StartTime float64 // 开始时间
	EndTime   float64 // 结束时间
	Total     uint    // 总次数
	Failed    uint    // 失败次数
	AvgTime   float32 // 平均用时
	MinTime   float32 // 最少用时
	MaxTime   float32 // 最多用时
}

type RequestLog struct {
	BaseLog
	App                string                 // 应用名
	Node               string                 // 处理请求的节点，ip:port
	ClientIp           string                 // 真实的用户IP，通过 X-Real-IP 续传
	FromApp            string                 // 调用方应用
	FromNode           string                 // 调用方节点，格式 ip:port
	ClientId           string                 // 客户唯一编号，通过 X-Client-ID 续传
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
	ResponseData       interface{}            // 指定要记录的数据内容，JSON对象，集合类型仅记录少量内容，敏感数据应脱敏，非对象内容不进行记录
}
