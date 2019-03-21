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

type LogLevelType int

const LogLevelDebug LogLevelType = 1
const LogLevelInfo LogLevelType = 2
const LogLevelWarning LogLevelType = 3
const LogLevelError LogLevelType = 4

type Logger struct {
	level  LogLevelType
	truncations []string
	writer func(string)
}

func (logger *Logger) SetLevel(level LogLevelType) {
	logger.level = level
}

func (logger *Logger) SetWriter(writer func(string)) {
	logger.writer = writer
}

func (logger *Logger) SetTruncations(truncations ...string) {
	logger.truncations = append(logger.truncations, truncations...)
}

func (logger *Logger) Debug(logType string, data ...interface{}) {
	logger.log(LogLevelDebug, logType, buildLogData(data...))
}

func (logger *Logger) Info(logType string, data ...interface{}) {
	logger.log(LogLevelInfo, logType, buildLogData(data...))
}

func (logger *Logger) Warning(logType string, data ...interface{}) {
	logger.trace(LogLevelWarning, logType, buildLogData(data...))
}

func (logger *Logger) Error(logType string, data ...interface{}) {
	logger.trace(LogLevelError, logType, buildLogData(data...))
}

func (logger *Logger) log(LogLevel LogLevelType, logType string, data map[string]interface{}) {
	settedLevel := logger.level
	if settedLevel == 0 {
		settedLevel = LogLevelInfo
	}
	if LogLevel < settedLevel {
		return
	}

	LogLevelName := standard.LogLevelInfo
	switch LogLevel {
	case LogLevelDebug:
		LogLevelName = standard.LogLevelDebug
	case LogLevelInfo:
		LogLevelName = standard.LogLevelInfo
	case LogLevelWarning:
		LogLevelName = standard.LogLevelWarning
	case LogLevelError:
		LogLevelName = standard.LogLevelError
	}

	data[standard.LogFieldLevel] = LogLevelName
	data[standard.LogFieldTime] = standard.MakeLogTime(time.Now())
	data[standard.LogFieldType] = logType
	buf, err := json.Marshal(data)

	if err != nil {
		// 无法序列化的数据包装为 JsonEncodeError
		buf, err = json.Marshal(map[string]interface{}{
			standard.LogFieldLevel: data[standard.LogFieldLevel],
			standard.LogFieldTime:  data[standard.LogFieldTime],
			standard.LogFieldType:  standard.LogTypeEncodingError,
			"data":                 fmt.Sprint(data),
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

func (logger *Logger) trace(LogLevel LogLevelType, logType string, data map[string]interface{}) {
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
	data[standard.LogFieldTraces] = strings.Join(traces, "; ")
	logger.log(LogLevel, logType, data)
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

func (logger *Logger) LogRequest(app, node, clientIp, fromApp, fromNode, clientId, sessionId, requestId, host string, authLevel, priority int, method, path string, requestHeaders map[string]string, requestData map[string]interface{}, usedTime float32, responseCode int, responseHeaders map[string]string, responseDataLength uint, responseData interface{}, extraInfo map[string]interface{}){
	extraInfo[standard.LogFieldRequestApp] = app
	extraInfo[standard.LogFieldRequestNode] = node
	extraInfo[standard.LogFieldRequestClientIp] = clientIp
	extraInfo[standard.LogFieldRequestFromApp] = fromApp
	extraInfo[standard.LogFieldRequestFromNode] = fromNode
	extraInfo[standard.LogFieldRequestClientId] = clientId
	extraInfo[standard.LogFieldRequestSessionId] = sessionId
	extraInfo[standard.LogFieldRequestRequestId] = requestId
	extraInfo[standard.LogFieldRequestHost] = host
	extraInfo[standard.LogFieldRequestAuthLevel] = authLevel
	extraInfo[standard.LogFieldRequestPriority] = priority
	extraInfo[standard.LogFieldRequestMethod] = method
	extraInfo[standard.LogFieldRequestPath] = path
	extraInfo[standard.LogFieldRequestRequestHeaders] = requestHeaders
	extraInfo[standard.LogFieldRequestArgs] = requestData
	extraInfo[standard.LogFieldRequestUsedTime] = usedTime
	extraInfo[standard.LogFieldRequestStatus] = responseCode
	extraInfo[standard.LogFieldRequestResponseHeaders] = responseHeaders
	extraInfo[standard.LogFieldRequestOutLen] = responseDataLength
	extraInfo[standard.LogFieldRequestResult] = responseData
	logger.log(LogLevelInfo, standard.LogTypeRequest, extraInfo)
}
