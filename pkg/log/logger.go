package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	LOG_TIME_FORMAT = "02-01-2006 15:04:05"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        LOG_TIME_FORMAT, // the "time" field configuration
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
	logLevel := os.Getenv("ENVIRONMENT")
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		lvl = logrus.DebugLevel
	}
	logrus.SetLevel(lvl)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	if len(arr) > 1 {
		if len(arr) > 2 {
			return strings.Join(arr[len(arr)-3:], "/")
		}
		return strings.Join(arr[len(arr)-2:], "/")
	}
	return arr[len(arr)-1]
}
