package api

import (
	"io"
	"net/http"
)

type ForwardRequest struct {
	Target string
	Method string
	Body   io.Reader
	Header http.Header
}

type ForwardResponse struct {
	TargetResponse *http.Response
}
