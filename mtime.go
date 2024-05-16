package mtime

import (
	"github.com/brahma-adshonor/gohook"
	_ "os"
	"sync/atomic"
	"time"
	_ "time"
	_ "unsafe"
)

var _offset_sec int64 = 0

func init() {
	err := gohook.Hook(time.Now, myTime, nil)
	if err != nil {
		panic(err)
	}
}

//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)
func nativeNow() time.Time {
	sec, nsec, _ := now()
	return time.Unix(sec, int64(nsec))
}
func myTime() time.Time {
	return nativeNow().Add(time.Duration(atomic.LoadInt64(&_offset_sec)))
}
func hookTime(sec time.Duration) {
	atomic.StoreInt64(&_offset_sec, int64(sec))
}

func SetMineTimeUnix(msec int64) {
	ResetTime()
	hookTime(time.UnixMilli(msec).Sub(time.Now()))
}

func SetMineTime(t time.Time) {
	ResetTime()
	hookTime(t.Sub(time.Now()))
}

func ResetTime() {
	atomic.StoreInt64(&_offset_sec, 0)
}
