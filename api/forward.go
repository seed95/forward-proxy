package api

import (
	"io"
	"net/http"
	"time"
)

type ForwardRequest struct {
	ReceivedAt time.Time
	Target     string // Target url
	Method     string
	Body       io.Reader
	Header     http.Header
}

type ForwardResponse struct {
	TargetResponse *http.Response
}
