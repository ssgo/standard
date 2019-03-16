package simple

import (
	"encoding/json"
	"fmt"
	"github.com/ssgo/standard"
	"log"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	level  standard.LogLevelType
	truncations []string
	writer func(string)
}

func (logger *Logger) SetLevel(level standard.LogLevelType) {
	logger.level = level
}

func (logger *Logger) SetWriter(writer func(string)) {
	logger.writer = writer
}

func (logger *Logger) SetTruncations(truncations ...string) {
	logger.truncations = append(logger.truncations, truncations...)
}

func (logger *Logger) Debug(logType string, data ...interface{}) {
	logger.log(standard.LogDebug, logType, buildLogData(data...))
}

func (logger *Logger) Info(logType string, data ...interface{}) {
	logger.log(standard.LogInfo, logType, buildLogData(data...))
}

func (logger *Logger) Warning(logType string, data ...interface{}) {
	logger.trace(standard.LogWarning, logType, buildLogData(data...))
}

func (logger *Logger) Error(logType string, data ...interface{}) {
	logger.trace(standard.LogError, logType, buildLogData(data...))
}

func (logger *Logger) log(logLevel standard.LogLevelType, logType string, data map[string]interface{}) {
	settedLevel := logger.level
	if settedLevel == 0 {
		settedLevel = standard.LogInfo
	}
	if logLevel < settedLevel {
		return
	}

	logLevelName := standard.LogInfoName
	switch logLevel {
	case standard.LogDebug:
		logLevelName = standard.LogDebugName
	case standard.LogInfo:
		logLevelName = standard.LogInfoName
	case standard.LogWarning:
		logLevelName = standard.LogWarningName
	case standard.LogError:
		logLevelName = standard.LogErrorName
	}

	data[standard.LogLevelFieldName] = logLevelName
	data[standard.LogTimeFieldName] = standard.MakeLogTime(time.Now())
	data[standard.LogTypeFieldName] = logType
	buf, err := json.Marshal(data)

	if err != nil {
		// 无法序列化的数据包装为 JsonEncodeError
		buf, err = json.Marshal(map[string]interface{}{
			standard.LogLevelFieldName: data[standard.LogLevelFieldName],
			standard.LogTimeFieldName:  data[standard.LogTimeFieldName],
			standard.LogTypeFieldName:  standard.LogEncodingErrorType,
			"data":                     fmt.Sprint(data),
		})
		return
	}

	if err == nil {
		if logger.writer == nil {
			log.Print(string(buf))
		} else {
			logger.writer(string(buf))
		}
	}
}

func (logger *Logger) trace(logLevel standard.LogLevelType, logType string, data map[string]interface{}) {
	traces := make([]string, 0)
	for i := 1; i < 20; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.Contains(file, "/go/src/") {
			continue
		}
		if logger.truncations != nil {
			for _, truncation := range logger.truncations {
				pos := strings.Index(file, truncation)
				if pos != -1 {
					file = file[pos+len(truncation):]
				}
			}
		}
		traces = append(traces, fmt.Sprintf("%s:%d", file, line))
	}
	data[standard.LogTracesFieldName] = strings.Join(traces, "; ")
	logger.log(logLevel, logType, data)
}

func buildLogData(args ...interface{}) map[string]interface{} {
	if len(args) == 1 {
		if mapData, ok := args[0].(map[string]interface{}); ok {
			return mapData
		}
	}
	data := map[string]interface{}{}
	for i:=1; i<len(args); i+=2 {
		if k, ok := args[i-1].(string); ok {
			data[k] = args[i]
		}
	}
	return data
}

func (logger *Logger) LogRequest(app, node, clientIp, caller, clientId, sessionId, requestId, host string, authLevel, priority int, method, path string, requestHeaders map[string]string, requestData map[string]interface{}, usedTime float32, responseCode int, responseHeaders map[string]string, responseDataLength uint, responseData interface{}, extraInfo map[string]interface{}){
	extraInfo[standard.LogRequestAppFieldName] = app
	extraInfo[standard.LogRequestNodeFieldName] = node
	extraInfo[standard.LogRequestClientIpFieldName] = clientIp
	extraInfo[standard.LogRequestCallerFieldName] = caller
	extraInfo[standard.LogRequestClientIdFieldName] = clientId
	extraInfo[standard.LogRequestSessionIdFieldName] = sessionId
	extraInfo[standard.LogRequestRequestIdFieldName] = requestId
	extraInfo[standard.LogRequestHostFieldName] = host
	extraInfo[standard.LogRequestAuthLevelFieldName] = authLevel
	extraInfo[standard.LogRequestPriorityFieldName] = priority
	extraInfo[standard.LogRequestMethodFieldName] = method
	extraInfo[standard.LogRequestPathFieldName] = path
	extraInfo[standard.LogRequestRequestHeadersFieldName] = requestHeaders
	extraInfo[standard.LogRequestArgsFieldName] = requestData
	extraInfo[standard.LogRequestUsedTimeFieldName] = usedTime
	extraInfo[standard.LogRequestStatusFieldName] = responseCode
	extraInfo[standard.LogRequestResponseHeadersFieldName] = responseHeaders
	extraInfo[standard.LogRequestOutLenFieldName] = responseDataLength
	extraInfo[standard.LogRequestResultFieldName] = responseData
	logger.log(standard.LogInfo, standard.LogRequestType, extraInfo)
}
