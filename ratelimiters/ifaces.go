package ratelimiters

import "time"

type RateLimiter interface {
	Allow(clientID string) bool
}

type Config struct {
	Duration     time.Duration
	RequestLimit int
}
