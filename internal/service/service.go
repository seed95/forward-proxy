package service

import (
	"context"
	"fmt"
	"github.com/seed95/forward-proxy/api"
	"github.com/seed95/forward-proxy/internal/repo"
	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"
	"net/http"
	"time"
)

type Service interface {
	// ForwardRequest forwards received requests to target URL
	ForwardRequest(ctx context.Context, req *api.ForwardRequest) (res *api.ForwardResponse, err error)

	// GetStats returns statistical information for received requests.
	GetStats(ctx context.Context, req *api.StatsRequest) (res *api.StatsResponse)
}

type service struct {
	client http.Client
	cache  repo.CacheRepo
}

var _ Service = (*service)(nil)

func New(cache repo.CacheRepo) Service {
	return &service{
		client: http.Client{},
		cache:  cache,
	}
}

func (s *service) ForwardRequest(ctx context.Context, req *api.ForwardRequest) (res *api.ForwardResponse, err error) {
	// Log request response
	defer func(startTime time.Time) {
		commonKeyVal := []keyval.Pair{
			keyval.String("req", fmt.Sprintf("%+v", req)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		log.ReqRes(startTime, err, commonKeyVal...)
	}(time.Now())

	// TODO check this segment
	cacheResp, err := s.cache.GetCachedRequest(ctx, req.Target)
	if err != nil {

	}
	_ = cacheResp

	proxyReq, err := http.NewRequest(req.Method, req.Target, req.Body)
	if err != nil {
		// TODO handle error
		return nil, err
	}
	proxyReq.Header = req.Header

	proxyRes, err := s.client.Do(proxyReq)
	if err != nil {
		// TODO handle error
		//http.Error(w, err.Error(), http.StatusBadGateway)
		return nil, err
	}

	if req.Method == http.MethodGet && proxyRes.StatusCode == http.StatusOK {
		go func() {
			if err := s.cache.CacheResponse(context.TODO(), req.Target, proxyRes); err != nil {
				// TODO log error
			}
		}()
	}

	res = &api.ForwardResponse{
		TargetResponse: proxyRes,
	}

	return res, nil
}

func (s *service) GetStats(ctx context.Context, req *api.StatsRequest) (res *api.StatsResponse) {
	// Log request response
	defer func(startTime time.Time) {
		commonKeyVal := []keyval.Pair{
			keyval.String("req", fmt.Sprintf("%+v", req)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		// TODO check error
		log.ReqRes(startTime, nil, commonKeyVal...)
	}(time.Now())

	// Make response
	res = &api.StatsResponse{}

	return res
}
