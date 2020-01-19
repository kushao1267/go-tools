package funcs

import "runtime"

type Logger interface {
	Error(args ...interface{})
}

var logger Logger

func CallFuncStack() string {
	buf := make([]byte, 1<<12) // 4096个字节
	return string(buf[:runtime.Stack(buf, false)])
}

func Try(fun func(), handler func(interface{})) { // recovery if panic occurred.
	defer func() {
		if err := recover(); err != nil {
			if handler == nil {
				logger.Error("panic catch:%v stack:%s", err, CallFuncStack())
			} else {
				handler(err)
			}
		}
	}()
	fun()
}