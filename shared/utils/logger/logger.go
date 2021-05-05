package logger

import (
	"fmt"
	"gitlab.com/dataon1/sf7-kit/shared/utils/config"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func MakeLogEntry(c echo.Context, fields map[string]interface{}) (*os.File, *logrus.Entry) {
	var (
		logFile *os.File
	)

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	if c == nil {
		if fields == nil {
			fields = make(map[string]interface{})
		}
		fields["app"] = config.GetString("name")
		fields["at"] = time.Now().Format("2006-01-02 15:04:05")
		return logFile, log.WithFields(fields)
	}

	return logFile, log.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
		"app":    config.GetString("name"),
	})
}


func Info(fields map[string]interface{}, args ...interface{}) {
	file, function, line := whereAmI(2)
	fields = caller(fields, file, function, line)
	logFile, LogEntry := MakeLogEntry(nil, fields)
	LogEntry.Info(args...)
	if logFile != nil {
		defer logFile.Close()
	}
}

func Debug(fields map[string]interface{}, args ...interface{}) {
	file, function, line := whereAmI(2)
	fields = caller(fields, file, function, line)
	logFile, LogEntry := MakeLogEntry(nil, fields)
	LogEntry.Debug(args...)
	if logFile != nil {
		defer logFile.Close()
	}
}

func Warn(fields map[string]interface{}, args ...interface{}) {
	file, function, line := whereAmI(2)
	fields = caller(fields, file, function, line)
	logFile, LogEntry := MakeLogEntry(nil, fields)
	LogEntry.Warn(args...)
	if logFile != nil {
		defer logFile.Close()
	}
}

func Error(fields map[string]interface{}, args ...interface{}) {
	file, function, line := whereAmI(2)
	fields = caller(fields, file, function, line)
	logFile, LogEntry := MakeLogEntry(nil, fields)
	LogEntry.Error(args...)
	if logFile != nil {
		defer logFile.Close()
	}
}

func Panic(fields map[string]interface{}, args ...interface{}) {
	file, function, line := whereAmI(2)
	fields = caller(fields, file, function, line)
	logFile, LogEntry := MakeLogEntry(nil, fields)
	LogEntry.Panic(args...)
	if logFile != nil {
		defer logFile.Close()
	}
}

func Fatal(fields map[string]interface{}, args ...interface{}) {
	file, function, line := whereAmI(2)
	fields = caller(fields, file, function, line)
	logFile, LogEntry := MakeLogEntry(nil, fields)
	LogEntry.Fatal(args...)
	if logFile != nil {
		defer logFile.Close()
	}
}

// return a string containing the file name, function name
// and the line number of a specified entry on the call stack
// https://github.com/jimlawless/whereami
func whereAmI(depthList ...int) (string, string, int) {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return chopPath(file), runtime.FuncForPC(function).Name(), line
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

// return info caller
func caller(fields map[string]interface{}, file string, function string, line int) map[string]interface{} {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["file"] = fmt.Sprintf("%s:%d", file, line)
	fields["function"] = function
	return fields
}

func logWithFile(log *logrus.Logger, logFile *os.File, cwd string) (*logrus.Logger, *os.File) {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to determine working directory: %s", err)
	}

	runID := config.GetAppName()
	time := time.Now().Format("2006-01-02")
	logLocation := filepath.Join(cwd+config.GetLogFile(), "log-"+runID+"-"+time+".log")
	logFile, err = os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("Failed to open log file %s for output: %s", logLocation, err)
	}

	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	log.SetOutput(logFile)

	return log, logFile
}
