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

const DiscoverDefaultRegistry = "127.0.0.1:6379::15" // 默认注册中心配置
const DiscoverHeaderClientIp = "X-Real-IP"           // 真实的用户IP，通过 X-Real-IP 续传
const DiscoverHeaderClientId = "X-Client-ID"         // 客户唯一编号，通过 X-Client-ID 续传
const DiscoverHeaderSessionId = "X-Session-ID"       // 会话唯一编号，通过 X-Session-ID 续传
const DiscoverHeaderRequestId = "X-Request-ID"       // 请求唯一编号，通过 X-Request-ID 续传
const DiscoverHeaderHost = "X-Host"                  // 真实用户请求的Host，通过 X-Host 续传
