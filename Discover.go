package standard

/*
	注册中心使用redis，存储类型 hash
	Key：app	应用名称
	Field：ip:port 节点地址
	Value：weight 权重

	订阅通道 CH_{app} CH_ 开头+应用名称

	注册中心配置格式 ip:port:password:db， 默认值 127.0.0.1:6379::15

	服务启动时在redis注册自己 HSET app ip:port weight
	并且向 CH_{app} 存入 PUBLISH CH_{app} "ip:port weight"

	服务停止时在redis注销自己 HDEL app ip:port
	并且向 CH_{app} 存入 PUBLISH CH_{app} "ip:port 0" 权重值为0表示注销

	服务间调用协议为 h2c，需要在每一次调用时传递下列头部信息以确保在每一个服务节点上能够有效处理和记录数据
*/

const DiscoverHeaderClientIp = "X-Real-IP"                    // 真实的用户IP
const DiscoverHeaderForwardedFor = "X-Forwarded-For"          // 客户端IP列表
const DiscoverHeaderClientId = "X-Client-ID"                  // 客户唯一编号
const DiscoverHeaderDeviceId = "X-Device-ID"                  // 设备唯一编号
const DiscoverHeaderClientAppName = "X-Client-App-Name"       // 客户端App名字
const DiscoverHeaderClientAppVersion = "X-Client-App-Version" // 客户端App版本号
const DiscoverHeaderSessionId = "X-Session-ID"                // 会话唯一编号
const DiscoverHeaderRequestId = "X-Request-ID"                // 请求唯一编号
const DiscoverHeaderHost = "X-Host"                           // 真实用户请求的Host
const DiscoverHeaderScheme = "X-Scheme"                       // 真实用户请求的 http or https
const DiscoverHeaderFromApp = "X-From-App"                    // 来源App
const DiscoverHeaderFromNode = "X-From-Node"                  // 来源节点
const DiscoverHeaderUserAgent = "X-User-Agent"                // 真实用户的UserAgent

var DiscoverRelayHeaders = []string{DiscoverHeaderClientIp, DiscoverHeaderForwardedFor, DiscoverHeaderClientId, DiscoverHeaderDeviceId, DiscoverHeaderClientAppName, DiscoverHeaderClientAppVersion, DiscoverHeaderSessionId, DiscoverHeaderRequestId, DiscoverHeaderHost, DiscoverHeaderScheme, DiscoverHeaderUserAgent}

//const DiscoverHeaderDeviceType = "X-Device-Type"
//const DiscoverHeaderDeviceName = "X-Device-Name"
//const DiscoverHeaderDeviceBrowserName = "X-Device-BrowserName"
//const DiscoverHeaderDeviceBrowserVersion = "X-Device-BrowserVersion"
//const DiscoverHeaderDeviceOSName = "X-Device-OSName"
//const DiscoverHeaderDeviceOSVersion = "X-Device-OSVersion"
//const DiscoverHeaderDeviceNetType = "X-Device-NetType"

const DiscoverDefaultRegistry = "127.0.0.1:6379::15" // 默认注册中心配置
const DiscoverEnvRegistry = "DISCOVER_REGISTRY"      // 注册中心地址，"127.0.0.1:6379:15"、"127.0.0.1:6379:15:password"
const DiscoverEnvApp = "DISCOVER_APP"                // 应用名，注册为服务
const DiscoverEnvWeight = "DISCOVER_WEIGHT"          // 应用权重，默认值：100
const DiscoverEnvCalls = "DISCOVER_CALLS"            // 被调用的应用定义，{"app1": "5000:token", "app2": "1000"}，也可以使用 DISCOVER_CALLS_app1=timeout:token，token会通过 Access-Token 头进行传递
