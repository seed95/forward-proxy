package repo

import (
	"context"
	"github.com/seed95/forward-proxy/internal/model"
)

type CacheRepo interface {
	// CacheResponse cache response received from target url
	CacheResponse(ctx context.Context, url string, res interface{}) (err error)

	// GetCachedRequest returns cached response for target url
	// If exist cache return response otherwise return error
	GetCachedRequest(ctx context.Context, url string) (res interface{}, err error)
}

type StatsRepo interface {
	// SaveStat save model.Statistical information for received request
	SaveStat(ctx context.Context, stat model.Statistical) (err error)

	// GetStats returns statistical information from (millisecond) variable onwards
	GetStats(ctx context.Context, from int64) (stats []model.Statistical, err error)
}
