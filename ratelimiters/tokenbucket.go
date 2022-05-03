package ratelimiters

import (
	"math"
	"sync"
	"time"
)

// Duration     time.Duration
// RequestLimit int
type tokenBucket struct {
	mu                  *sync.Mutex
	maxTokens           int64
	config              Config
	currTokens          int64
	lastRefillTimestamp time.Time
}

func NewTokenBucket(config Config) *tokenBucket {
	return &tokenBucket{
		maxTokens:           config.RequestLimit,
		lastRefillTimestamp: time.Now(),
		config:              config,
	}
}

func (t *tokenBucket) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.refill()
	if t.currTokens > 0 {
		t.currTokens = t.currTokens - 1
		return true
	}
	return false
}

func (t *tokenBucket) refill() {
	now := time.Now()
	end := time.Since(t.lastRefillTimestamp)
	tokensToAdd := t.tokens(end)
	t.currTokens = int64(math.Min(float64(t.currTokens)+tokensToAdd, float64(t.maxTokens)))
	t.lastRefillTimestamp = now
}

func (t *tokenBucket) tokens(end time.Duration) float64 {
	return (end.Seconds() / t.config.Duration.Seconds()) * float64(t.config.RequestLimit)
}

func insert(intervals [][]int, newInterval []int) [][]int {
    var output [][]int 
    var i int
    _, end := newInterval[0], newInterval[1]
    for ;i < len(intervals);i++ {
        if intervals[i][1] < end {
            output = append(output, intervals[i])
        }
    }

    output = append(output, newInterval)
    for ;i < len(intervals); i++ {
        if len(output) == 0 {
            output = append(output, intervals[i])
            continue
        }
        prev := output[len(output)-1]
        curr := intervals[i] // (5 9)
        if prev[1] >= curr[0] { 
            output = output[:len(output)-1] // output.pop()
            output = append(output, []int{prev[0], max(prev[1], curr[1])}) // (1, 9)
            
        } else {
            output = append(output, curr)
        }
    }
    return output
}

func max(a, b int) int {
    if a > b { return a }
    return b
}