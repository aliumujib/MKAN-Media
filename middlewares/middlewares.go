package mws

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

type exception struct {
	Message string `json:"message"`
}

//Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

//AccessLogToConsole prints sever logs to the terminal
func AccessLogToConsole(r http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		logger := handlers.CombinedLoggingHandler(os.Stdout, r)

		logger.ServeHTTP(w, req)

	})
}
