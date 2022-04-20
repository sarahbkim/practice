package main

import "testing"

func Test_rateLimiter_isAllowed(t *testing.T) {
	type fields struct {
		clientIDs    map[string][]int64
		maxReqs      int
		intervalSecs int
	}
	type args struct {
		clientID string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     bool
		prevCall int
	}{
		{
			name: "",
			fields: fields{
				clientIDs:    map[string][]int64{},
				maxReqs:      3,
				intervalSecs: 30,
			},
			args: args{
				clientID: "client1",
			},
			want:     true,
			prevCall: 0,
		},
		{
			name: "",
			fields: fields{
				clientIDs:    map[string][]int64{},
				maxReqs:      3,
				intervalSecs: 30,
			},
			args: args{
				clientID: "client1",
			},
			want:     true,
			prevCall: 1,
		},
		{
			name: "",
			fields: fields{
				clientIDs:    map[string][]int64{},
				maxReqs:      3,
				intervalSecs: 30,
			},
			args: args{
				clientID: "client1",
			},
			want:     false,
			prevCall: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rateLimiter{
				clientIDs:    tt.fields.clientIDs,
				maxReqs:      tt.fields.maxReqs,
				intervalSecs: tt.fields.intervalSecs,
			}
			for i := 0; i < tt.prevCall; i++ {
				r.isAllowed(tt.args.clientID)
			}
			if got := r.isAllowed(tt.args.clientID); got != tt.want {
				t.Errorf("rateLimiter.isAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}
