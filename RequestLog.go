package standard

const LogTypeRequest = "request" // 服务请求日志类型

const LogFieldRequestApp = "app"                         // 应用名
const LogFieldRequestNode = "node"                       // 处理请求的节点，ip:port
const LogFieldRequestClientIp = "clientIp"               // 真实的用户IP，通过 X-Real-IP 续传
const LogFieldRequestFromApp = "fromApp"                 // 调用方应用
const LogFieldRequestFromNode = "fromNode"               // 调用方节点，格式 ip:port
const LogFieldRequestClientId = "clientId"               // 客户唯一编号，通过 X-Client-ID 续传
const LogFieldRequestSessionId = "sessionId"             // 会话唯一编号，通过 X-Session-ID 续传
const LogFieldRequestRequestId = "requestId"             // 请求唯一编号，通过 X-Request-ID 续传
const LogFieldRequestHost = "host"                       // 真实用户请求的Host，通过 X-Host 续传
const LogFieldRequestScheme = "scheme"                   // http scheme, http or https
const LogFieldRequestProto = "proto"                     // http proto, 1.1 or 2.0
const LogFieldRequestAuthLevel = "authLevel"             // 验证级别，用来校验用户是否有权限访问
const LogFieldRequestPriority = "priority"               // 优先级，用来在服务故障时进行自动降级处理
const LogFieldRequestMethod = "method"                   // 请求的方法
const LogFieldRequestPath = "path"                       // 请求的路径，不包括GET参数部分，如果有PATH参数应该记录定义的PATH
const LogFieldRequestRequestHeaders = "requestHeaders"   // 请求头，排除掉指定不需要信息后的所有头部内容，敏感数据应脱敏
const LogFieldRequestArgs = "requestData"                // 请求的数据内容，JSON对象，集合类型仅记录少量内容，敏感数据应脱敏，非对象内容过大应做截取
const LogFieldRequestUsedTime = "usedTime"               // 处理请求花费的时间，格式为float32，单位毫秒
const LogFieldRequestStatus = "responseCode"             // 应答代码，200 1000+ 正常应答，201～399，1～199  600～999 特殊应答，<1 异常应答
const LogFieldRequestResponseHeaders = "responseHeaders" // 应答头，排除掉指定不需要信息后的所有头部内容，敏感数据应脱敏
const LogFieldRequestOutLen = "responseDataLength"       // 应答的数据长度
const LogFieldRequestResult = "responseData"             // 指定要记录的数据内容，JSON对象，集合类型仅记录少量内容，敏感数据应脱敏，非对象内容不进行记录
