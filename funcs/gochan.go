package funcs

import (
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var (
	stopFromSys = make(chan os.Signal, 1)
	stopSignal  = make(chan bool)
	isRunning   = int32(1)
)

func Stop() {
	if atomic.CompareAndSwapInt32(&isRunning, 1, 0) {
		close(stopSignal)
	}
}

// WaitForSystemExit 接收系统Signal
func WaitForSystemExit() {
	signal.Notify(stopFromSys, os.Interrupt, os.Kill, syscall.SIGTERM)
	select {
	case <-stopFromSys:
		Stop()
	}
}

// Go 拥有recovery能力的goroutine
func Go(fn func()) {
	go func() {
		Try(fn, nil)
	}()
}

// GoChan 监听各种系统Signal且拥有recovery能力的goroutine
func GoChan(fn func(cstop <-chan bool)) {
	Go(func() {
		fn(stopSignal)
	})
}
