package ratelimiters

import (
	"sync"
	"time"
)

type leakyBucket struct {
	config Config
	count  int64
	mu     *sync.Mutex
}

func NewLeakyBucket(done <-chan struct{}, config Config) *leakyBucket {
	l := &leakyBucket{
		config: config,
		count:  0,
		mu:     &sync.Mutex{},
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
			l.count = l.config.RequestLimit
			l.mu.Unlock()
		}
	}
}

func (l *leakyBucket) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.count == 0 {
		return false
	}
	l.count--
	return true
}
