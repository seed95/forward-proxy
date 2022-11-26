package handler

import (
	"fmt"
	"net/http"

	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"
)

// logMiddleware Log request
func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("http_request", keyval.String("req", fmt.Sprintf("%+v", r)))
		next(w, r)
	}
}
