package ratelimiters

import (
	"sync"
	"time"
)

type leakyBucket struct {
	config           Config
	requestsByClient map[string]int
	mu               *sync.Mutex
}

func NewLeakyBucket(done <-chan struct{}, config Config) RateLimiter {
	l := &leakyBucket{
		config:           config,
		requestsByClient: map[string]int{},
		mu:               &sync.Mutex{},
	}
	go l.reset(done)
	return l
}

func (l *leakyBucket) reset(done <-chan struct{}) {
	ticker := time.NewTicker(l.config.Duration)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			l.mu.Lock()
			for c := range l.requestsByClient {
				l.requestsByClient[c] = l.config.RequestLimit
			}
			l.mu.Unlock()
		}
	}
}

func (l *leakyBucket) Allow(clientID string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	counts, ok := l.requestsByClient[clientID]
	if !ok {
		l.requestsByClient[clientID] = l.config.RequestLimit - 1
		return true
	}
	if counts == 0 {
		return false
	}
	l.requestsByClient[clientID]--
	return true
}
