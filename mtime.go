package mtime

import (
	"github.com/brahma-adshonor/gohook"
	_ "os"
	"time"
	_ "time"
	_ "unsafe"
)

var _offset_sec time.Duration = 0

//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)
func nativeNow() time.Time {
	sec, nsec, _ := now()
	return time.Unix(sec, int64(nsec))
}
func myTime() time.Time {
	return nativeNow().Add(_offset_sec)
}
func hookTime(sec time.Duration) error {
	_offset_sec = sec
	return gohook.Hook(time.Now, myTime, nil)
}

func SetMineTimeUnix(msec int64) error {
	return hookTime(-time.Since(time.UnixMilli(msec)))
}

func SetMineTime(t time.Time) error {
	return hookTime(t.Sub(time.Now()))
}
