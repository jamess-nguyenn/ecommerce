package middleware

import (
	"ecommerce/helpers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

const LogGroupHandler = "http_handlers"

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.ResetTrackingIDs(r.Header)
		helpers.LogInfo("Incoming HTTP request", getLogHandler(r), getAdditionalData(r))

		next.ServeHTTP(w, r)
	})
}

func getLogHandler(r *http.Request) map[string]any {
	return map[string]any{
		"Method":        r.Method,
		"Header":        r.Header,
		"Params":        helpers.GetParams(r),
		"Body":          r.Body,
		"ContentLength": r.ContentLength,
		"Host":          r.Host,
		"RequestURI":    r.RequestURI,
		"Response":      r.Response,
	}
}

func getAdditionalData(r *http.Request) map[string]any {
	function := mux.CurrentRoute(r).GetName()

	handlerName := getHandler(r.RequestURI)
	group := getLogGroup(handlerName)

	return map[string]any{
		"group":    group,
		"function": function,
	}
}

func getHandler(requestURI string) string {
	requestURI = strings.Trim(requestURI, "/")

	if parts := strings.Split(requestURI, "/"); len(parts) > 1 {
		return strings.TrimSuffix(parts[1], "s")
	}

	return ""
}

func getLogGroup(handlerName string) string {
	if handlerName != "" {
		return fmt.Sprintf("%s_%s", LogGroupHandler, handlerName)
	}

	return LogGroupHandler
}
