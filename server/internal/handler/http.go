package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/seed95/forward-proxy/api"
	"github.com/seed95/forward-proxy/internal/service"
	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"

	"golang.org/x/time/rate"
)

type httpHandler struct {
	srv     service.Service
	limiter *rate.Limiter
}

func NewHttpHandler(srv service.Service, limiter *rate.Limiter) *httpHandler {
	handler := &httpHandler{
		srv:     srv,
		limiter: limiter,
	}
	return handler
}

func (h *httpHandler) Serve(w http.ResponseWriter, r *http.Request) {
	// Regex for stats route
	reg := regexp.MustCompile(`/stats\?time=\d`)
	switch {
	case reg.MatchString(r.URL.String()):
		// Returns proxy statistical information
		h.getStat(w, r)
	default:
		// Forward received requests
		h.forward(w, r)
	}
}

func (h *httpHandler) forward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Checking that the client does not send too many requests
	if !h.limiter.Allow() {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

	// Remove first slash from request
	targetUrl := r.URL.String()
	if strings.HasPrefix(r.URL.String(), "/") {
		targetUrl = strings.Replace(targetUrl, "/", "", 1)
	}

	// Make service request
	forwardReq := api.ForwardRequest{
		ReceivedAt: time.Now(),
		TargetUrl:  targetUrl,
		Method:     r.Method,
		Body:       r.Body,
		Header:     r.Header,
	}

	// Call service
	res, err := h.srv.ForwardRequest(r.Context(), &forwardReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(res.StatusCode)
	if _, err = w.Write(res.Body); err != nil {
		log.Error("write response", keyval.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *httpHandler) getStat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Make service request
	statsReq := api.StatsRequest{From: r.URL.Query().Get("time")}

	// Call service
	res, err := h.srv.GetStats(r.Context(), &statsReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	_ = json.NewEncoder(w).Encode(res)
}
