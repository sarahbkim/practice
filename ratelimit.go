package main

import "time"

type rateLimiter struct {
	clientIDs    map[string][]int64
	maxReqs      int
	intervalSecs int
}

func (r *rateLimiter) isAllowed(clientID string) bool {
	now := time.Now().Unix()
	if reqs, ok := r.clientIDs[clientID]; ok {
		if len(reqs) >= r.maxReqs {
			timeDiff := now - reqs[len(reqs)-1]
			reqs = reqs[1:]
			r.clientIDs[clientID] = append(reqs, now)
			if timeDiff < int64(r.intervalSecs) {
				return false
			}
		} else {
			r.clientIDs[clientID] = append(reqs, now)
		}
	} else {
		r.clientIDs[clientID] = []int64{now}
	}
	return true
}
