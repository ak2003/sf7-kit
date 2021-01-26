package logger

import (
	"os"

	"github.com/go-kit/kit/log"
)

func MakeLogEntry(service string, method string) log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", service,
			"method", method,
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	return logger
}

//func Info(args ...interface{}) {
//	LogEntry := MakeLogEntry()
//	err := level.Info(LogEntry).Log(nil, args )
//	if err != nil {
//		fmt.Errorf("write log info error %+v", err)
//	}
//}