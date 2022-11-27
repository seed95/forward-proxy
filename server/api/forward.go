package api

import (
	"io"
	"net/http"
	"time"
)

type ForwardRequest struct {
	ReceivedAt time.Time
	TargetUrl  string
	Method     string
	Body       io.Reader
	Header     http.Header
}

type ForwardResponse struct {
	StatusCode int
	Body       []byte
}
