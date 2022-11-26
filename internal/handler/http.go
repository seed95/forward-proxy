package handler

import (
	"encoding/json"
	"fmt"
	"github.com/seed95/forward-proxy/api"
	"github.com/seed95/forward-proxy/internal/service"
	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"
	"net/http"
	"regexp"
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
	h.mux.HandleFunc("/", h.serve)
	return h.mux
}

func (h *httpHandler) serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RequestURI())
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
	// TODO check postman body response
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "application/text")

	//if strings.Contains(r.URL.String(),"/stats?time=")

	//fmt.Println("req url:", r.URL)
	//fmt.Println("req url schema:", r.URL.Scheme)
	//fmt.Println("req raw url:", r.URL.RawPath)
	//fmt.Println("req raw query:", r.URL.RawQuery)
	//fmt.Println("req url string:", r.URL.String())
	//fmt.Println("req url path:", r.URL.Path)
	//fmt.Println("req url host:", r.URL.Host)
	//fmt.Println("req host:", r.Host)
	//fmt.Println("req remote addr:", r.RemoteAddr)
	//fmt.Println("req parse form:", r.ParseForm())
	//fmt.Println("req form:", r.Form)
	//fmt.Println("req", r)

	forwardReq := api.ForwardRequest{
		ReceivedAt: time.Now(),
		//Target:     r.URL.String(),
		TargetUrl: "https://www.google.com", // TODO check
		Method:    r.Method,
		Body:      r.Body,
		Header:    r.Header,
	}

	res, err := h.srv.ForwardRequest(r.Context(), &forwardReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(res.Body); err != nil {
		log.Error("write response", keyval.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(res.StatusCode)
}

func (h *httpHandler) getStat(w http.ResponseWriter, r *http.Request) {
	statsReq := api.StatsRequest{From: r.URL.Query().Get("time")}

	res, err := h.srv.GetStats(r.Context(), &statsReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}
