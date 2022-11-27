package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/seed95/forward-proxy/api"
	"github.com/seed95/forward-proxy/internal/model"
	"github.com/seed95/forward-proxy/internal/repo"
	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"
)

type Service interface {
	// ForwardRequest forwards received requests to target URL
	ForwardRequest(ctx context.Context, req *api.ForwardRequest) (res *api.ForwardResponse, err error)

	// GetStats returns statistical information for received requests.
	GetStats(ctx context.Context, req *api.StatsRequest) (res *api.StatsResponse, err error)
}

type service struct {
	client    http.Client
	cache     repo.CacheRepo
	statsRepo repo.StatsRepo
}

var _ Service = (*service)(nil)

func New(cache repo.CacheRepo, statsRepo repo.StatsRepo) Service {
	return &service{
		client:    http.Client{},
		cache:     cache,
		statsRepo: statsRepo,
	}
}

func (s *service) ForwardRequest(ctx context.Context, req *api.ForwardRequest) (res *api.ForwardResponse, err error) {
	// Log request response
	defer func(startTime time.Time) {
		commonKeyVal := []keyval.Pair{
			keyval.String("req", fmt.Sprintf("%+v", req)),
		}
		if res != nil {
			commonKeyVal = append(commonKeyVal, keyval.String("res_status_code", fmt.Sprintf("%v", res.StatusCode)))
			commonKeyVal = append(commonKeyVal, keyval.String("res_length", fmt.Sprintf("%v", len(res.Body))))
		}
		log.ReqRes(startTime, err, commonKeyVal...)
	}(time.Now())

	if cacheRes := s.cache.GetCachedRequest(ctx, req.TargetUrl); cacheRes != nil {
		return &api.ForwardResponse{
			StatusCode: http.StatusOK,
			Body:       cacheRes,
		}, nil
	}

	proxyReq, err := http.NewRequestWithContext(ctx, req.Method, req.TargetUrl, req.Body)
	if err != nil {
		return nil, err
	}
	proxyReq.Header = req.Header

	proxyRes, err := s.client.Do(proxyReq)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(proxyRes.Body)
	if err != nil {
		return nil, err
	}

	// Save statistical and cache response
	if req.Method == http.MethodGet {
		// Save statistical information in database
		stat := model.Statistical{
			Url:                  req.TargetUrl,
			StatusCode:           proxyRes.StatusCode,
			DurationResponseTime: time.Since(req.ReceivedAt).Milliseconds(),
			ReceivedAt:           req.ReceivedAt.UnixMilli(),
		}
		go func(stat model.Statistical) {
			if err := s.statsRepo.SaveStat(context.TODO(), stat); err != nil {
				log.Error("save stat", keyval.Error(err))
			}
		}(stat)

		// Cache response if status code is http.StatusOK
		if proxyRes.StatusCode == http.StatusOK {
			go func(targetUrl string, body []byte) {
				if err := s.cache.CacheResponse(context.TODO(), targetUrl, body); err != nil {
					log.Error("cache response", keyval.Error(err))
				}
			}(req.TargetUrl, body)
		}
	}

	res = &api.ForwardResponse{
		StatusCode: proxyRes.StatusCode,
		Body:       body,
	}

	return res, nil
}

func (s *service) GetStats(ctx context.Context, req *api.StatsRequest) (res *api.StatsResponse, err error) {
	// Log request response
	defer func(startTime time.Time) {
		commonKeyVal := []keyval.Pair{
			keyval.String("req", fmt.Sprintf("%+v", req)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		log.ReqRes(startTime, err, commonKeyVal...)
	}(time.Now())

	// Convert string from to duration
	from, err := time.ParseDuration(req.From + "m")
	duration := time.Now().UnixMilli() - from.Milliseconds()

	// Get stats
	stats, err := s.statsRepo.GetStats(ctx, duration)
	if err != nil {
		return nil, err
	}

	nSuccess := 0
	nFailed := 0
	durations := api.MaxDurations{}
	for _, stat := range stats {
		if stat.StatusCode == http.StatusOK {
			nSuccess++
		} else {
			nFailed++
		}

		if durations[stat.Url] < stat.DurationResponseTime {
			durations[stat.Url] = stat.DurationResponseTime
		}
	}

	// Make response
	res = &api.StatsResponse{
		ForwardingStats: api.ForwardingStats{
			SuccessCount: nSuccess,
			FailCount:    nFailed,
		},
		MaxDurations: durations,
	}
	return res, nil
}
