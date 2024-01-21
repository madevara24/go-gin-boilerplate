package logger

import (
	"fmt"
	"runtime"
)

func GetCallerTrace() Field {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return Field{Key: "source", Val: fmt.Sprintf("[%s:%d]", file, line)}
	}
	return Field{}
}
