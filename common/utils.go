package common

import (
	"fmt"
	"runtime"
)

func SendPanic(status int, errorString string)  {
	data := &ErrorData{
		ErrorString: errorString + GetFileInfo(2),
		Status:      status,
	}
	panic(data)
}

func GetFileInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return " [get file info failed] "
	}
	fn := runtime.FuncForPC(pc)
	return fmt.Sprintf(" [%s %s:%d] ", fn.Name(), file, line)
}