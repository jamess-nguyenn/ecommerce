package helpers

import (
	"ecommerce/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	RequestIDLength = 20
	TraceIDLength   = 32
)

var entry *logrus.Entry

func init() {
	log := logrus.New()

	// Set log level and formatter
	log.SetLevel(GetLogLevel())
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "datetime",
		},
	})

	// Configure log output file
	logPath := fmt.Sprintf(config.Logging.OutputFile, GetDate())
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(file)

	// Initialize entry with fields that should be in every log
	entry = log.WithFields(getOriginalFields())
}

func ResetTrackingIDs(h http.Header) {
	entry = entry.WithFields(getTrackingIDs(h))
}

func getOriginalFields() logrus.Fields {
	fields := getTrackingIDs(http.Header{})

	environment := GetEnv()
	fields["environment"] = environment

	return fields
}

func getTrackingIDs(h http.Header) logrus.Fields {
	requestID := GenerateRequestID()

	traceID := h.Get("X-Trace-ID")
	if traceID == "" {
		traceID = GenerateTraceID()
	}

	return logrus.Fields{
		"request_id": requestID,
		"trace_id":   traceID,
	}
}

func GetLogLevel() logrus.Level {
	if level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL")); err == nil {
		return level
	}

	return logrus.InfoLevel
}

func Log(level string, message string, data ...map[string]any) {
	dataField, additionalData := parseLogData(data, getAdditionalData())

	contextEntry := entry.WithFields(getAdditionalFields(dataField, additionalData))

	switch strings.ToLower(level) {
	case "debug":
		contextEntry.Debug(message)
	case "warning":
		contextEntry.Warning(message)
	case "error":
		contextEntry.Error(message)
	default:
		contextEntry.Info(message) // Default to Info if level is unrecognized
	}
}

func parseLogData(data []map[string]any, additionalData map[string]any) (map[string]any, map[string]any) {
	switch len(data) {
	case 1:
		return data[0], additionalData
	case 2:
		return data[0], data[1]
	default:
		return map[string]any{}, additionalData
	}
}

func LogDebug(message string, data ...map[string]any) {
	Log(getLevelName(), message, data...)
}

func LogInfo(message string, data ...map[string]any) {
	Log(getLevelName(), message, data...)
}

func LogWarning(message string, data ...map[string]any) {
	Log(getLevelName(), message, data...)
}

func LogError(message string, data ...map[string]any) {
	Log(getLevelName(), message, data...)
}

func getLevelName() string {
	functionInfo, _ := GetTrace(2)

	function := GetFunctionAfterDot(functionInfo)

	return strings.TrimPrefix(function, "Log")
}

func GenerateRequestID() string {
	return GenerateString(RequestIDLength)
}

func GenerateTraceID() string {
	return GenerateString(TraceIDLength)
}

func getAdditionalFields(data, additionalData map[string]any) logrus.Fields {
	fields := logrus.Fields{
		"data": data,
	}

	for key, value := range additionalData {
		fields[key] = value
	}

	return fields
}

func getAdditionalData() map[string]any {
	functionInfo, filePath := GetTrace(4)

	function := GetFunctionAfterDot(functionInfo)

	rootPath, _ := os.Getwd()
	group := getLogGroup(filePath, rootPath)

	return map[string]any{
		"group":    group,
		"function": function,
	}
}

func getLogGroup(filePath, rootPath string) string {
	group := strings.TrimPrefix(filePath, rootPath+"/")
	group = strings.TrimSuffix(group, ".go")
	group = strings.ReplaceAll(group, "/", "_")

	return group
}
