package mtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSetMineTimeUnix(t *testing.T) {
	type args struct {
		msec int64
	}
	msec1 := time.Now().Add(24 * time.Hour).UnixMilli()
	msec2 := time.Now().Add(-24 * time.Hour).UnixMilli()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "neg",
			args: args{msec: msec2},
		},
		{
			name: "pos",
			args: args{msec: msec1},
		},
	}
	t.Logf("tests=%+v", tests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMineTimeUnix(tt.args.msec)
			assert.Equal(t, tt.args.msec/1000, time.Now().Unix())
		})
	}
}

func TestSetMineTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	t1 := time.Now().Add(24 * time.Hour)
	t2 := time.Now().Add(-24 * time.Hour)
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "pos",
			args: args{t: t1},
		},
		{
			name: "neg",
			args: args{t: t2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMineTime(tt.args.t)
			assert.Equal(t, tt.args.t.Unix(), time.Now().Unix())
		})
	}
}
