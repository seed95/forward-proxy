package repo

import (
	"context"
	"github.com/seed95/forward-proxy/internal/model"
)

type CacheRepo interface {
	// CacheResponse cache received response body from target url
	CacheResponse(ctx context.Context, url string, body []byte) (err error)

	// GetCachedRequest returns cached response body for target url
	// If exist cache return body otherwise return nil
	GetCachedRequest(ctx context.Context, url string) (body []byte)
}

type StatsRepo interface {
	// SaveStat save model.Statistical information for received request
	SaveStat(ctx context.Context, stat model.Statistical) (err error)

	// GetStats returns statistical information from (millisecond) variable onwards
	GetStats(ctx context.Context, from int64) (stats []model.Statistical, err error)
}
