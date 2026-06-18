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
		"URL":           r.Host + r.URL.Path,
		"Headers":       r.Header,
		"Body":          r.Body,
		"ContentLength": r.ContentLength,
		"PathParams":    helpers.GetPathParams(r),
		"QueryParams":   r.URL.Query(),
		"Response":      r.Response,
	}
}

func getAdditionalData(r *http.Request) map[string]any {
	function := mux.CurrentRoute(r).GetName()

	path := getHandler(r.URL.Path)
	group := getLogGroup(path)

	return map[string]any{
		"group":    group,
		"function": function,
	}
}

func getHandler(path string) string {
	path = strings.Trim(path, "/")

	if parts := strings.Split(path, "/"); len(parts) > 2 {
		return strings.TrimSuffix(parts[2], "s")
	}

	return ""
}

func getLogGroup(handlerName string) string {
	if handlerName != "" {
		return fmt.Sprintf("%s_%s", LogGroupHandler, handlerName)
	}

	return LogGroupHandler
}
