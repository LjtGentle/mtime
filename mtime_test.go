package mtime

import (
	"fmt"
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
			name:    "pos",
			args:    args{msec: msec1},
			wantErr: false,
		},
		{
			name:    "neg",
			args:    args{msec: msec2},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetMineTimeUnix(tt.args.msec)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetMineTimeUnix() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, time.Now().Unix(), tt.args.msec/1000)
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
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:    "pos",
			args:    args{t: t1},
			wantErr: nil,
		},
		{
			name:    "neg",
			args:    args{t: t2},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, SetMineTime(tt.args.t), fmt.Sprintf("SetMineTime(%v)", tt.args.t))
			assert.Equal(t, tt.args.t.Unix(), time.Now().Unix())
		})
	}
}
