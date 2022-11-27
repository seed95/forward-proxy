package handler

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"
)

// LogMiddleware Log request
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("http_request", keyval.String("req", fmt.Sprintf("%+v", r)))
		next(w, r)
	}
}

// RecoverMiddleware recover panic as return error
func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Recover panic
		defer func() {
			if r := recover(); r != nil {
				stack := string(debug.Stack())
				log.Error("panic", keyval.String("stacktrace", stack))
			}
		}()

		next(w, r)
	}
}
