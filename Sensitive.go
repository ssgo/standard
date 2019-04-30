package standard

const SensitiveENVKey = "LOG_SENSITIVE"           // 敏感信息字段，[]string 的json格式
const RegexSensitiveENVKey = "LOG_REGEXSENSITIVE" // 敏感信息正则表达式，[]string 的json格式
const SensitiveRuleENVKey = "LOG_SENSITIVERULE"   // 敏感处理规则，[]string 的json格式，格式：threshold:leftNum*rightNum，例如：11:3*4 表示大于12字节的保留前3位和后4位
