package repo

import "context"

type CacheRepo interface {
	// CacheResponse cache response received from target url
	CacheResponse(ctx context.Context, url string, res interface{}) (err error)

	// GetCachedRequest returns cached response for target url
	// If exist cache return response otherwise return error
	GetCachedRequest(ctx context.Context, url string) (res interface{}, err error)
}

type StatRepo interface {
}
