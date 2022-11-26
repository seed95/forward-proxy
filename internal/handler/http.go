package handler

import (
	"fmt"
	"github.com/seed95/forward-proxy/api"
	"github.com/seed95/forward-proxy/internal/service"
	"net/http"
	"time"
)

type httpHandler struct {
	mux *http.ServeMux
	srv service.Service
}

func NewHttpHandler(srv service.Service) *httpHandler {
	handler := &httpHandler{
		mux: http.NewServeMux(),
		srv: srv,
	}
	return handler
}

func (h *httpHandler) Route() http.Handler {
	// Forward received requests
	h.mux.HandleFunc("/", h.forward)

	// Returns proxy statistical information
	h.mux.HandleFunc("/stats/:time", h.getStat)

	return h.mux
}

func (h *httpHandler) forward(w http.ResponseWriter, r *http.Request) {
	fmt.Println("req url:", r.URL)
	fmt.Println("req raw url:", r.URL.RawPath)

	forwardReq := api.ForwardRequest{
		ReceivedAt: time.Now(),
		Target:     r.URL.String(),
		Method:     r.Method,
		Body:       r.Body,
		Header:     r.Header,
	}

	res, err := h.srv.ForwardRequest(r.Context(), &forwardReq)
	if err != nil {
		// TODO handle error
	}
	_ = res
}

func (h *httpHandler) getStat(w http.ResponseWriter, r *http.Request) {
	strTime := r.URL.Query().Get("time")
	from, err := time.ParseDuration(strTime + "m")
	if err != nil {
		//TODO handle error
	}
	statsReq := api.StatsRequest{From: from}

	res := h.srv.GetStats(r.Context(), &statsReq)
	_ = res

}
