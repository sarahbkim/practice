package ratelimiters

import (
	"sync"
	"testing"
	"time"
)

func Test_leakyBucket_Allow(t *testing.T) {
	type fields struct {
		config           Config
		requestsByClient map[string]int
		nReqs            int
		every            time.Duration
	}
	type args struct {
		clientID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		run    func(l *leakyBucket, nreqs int, every time.Duration)
	}{
		{
			name: "0",
			fields: fields{
				config: Config{
					RequestLimit: 3,
					Duration:     time.Second * 3,
				},
				requestsByClient: map[string]int{},
				nReqs:            1,
				every:            1 * time.Second,
			},
			run: func(l *leakyBucket, nreqs int, every time.Duration) {
				ticker := time.NewTicker(every)
				for range ticker.C {
					l.Allow("client_id")
					nreqs--
					if nreqs == 0 {
						return
					}
				}
			},
			args: args{
				clientID: "client_id",
			},
			want: true,
		},
		{
			name: "1",
			fields: fields{
				config: Config{
					RequestLimit: 3,
					Duration:     time.Second * 3,
				},
				requestsByClient: map[string]int{},
				nReqs:            3,
				every:            1 * time.Second,
			},
			args: args{
				clientID: "client_id",
			},
			run: func(l *leakyBucket, nreqs int, every time.Duration) {
				ticker := time.NewTicker(every)
				for range ticker.C {
					l.Allow("client_id")
					nreqs--
					if nreqs == 0 {
						return
					}
				}
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &leakyBucket{
				config:           tt.fields.config,
				requestsByClient: tt.fields.requestsByClient,
				mu:               &sync.Mutex{},
			}
			tt.run(l, tt.fields.nReqs, tt.fields.every)
			if got := l.Allow(tt.args.clientID); got != tt.want {
				t.Errorf("leakyBucket.Allow() = %v, want %v", got, tt.want)
			}
		})
	}
}
