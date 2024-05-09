package gzlog

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
)

type Logger struct {
	LogLevel      gormlogger.LogLevel
	SlowThreshold time.Duration
}

func New(level gormlogger.LogLevel) Logger {
	return Logger{
		LogLevel:      level,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l Logger) SetAsDefault() {
	gormlogger.Default = l
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return Logger{
		SlowThreshold: l.SlowThreshold,
		LogLevel:      level,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}

	logrus.Infof("", args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}

	logrus.Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}

	logrus.Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error:
		sql, _ := fc()
		logrus.Error("[TRACE] ", err, " - sql query: ", sql)
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, _ := fc()
		logrus.Warn("[TRACE] ", " - sql query: ", sql)
	case l.LogLevel >= gormlogger.Info:
		sql, _ := fc()
		logrus.Info("[TRACE] ", " - sql query: ", sql)
	}
}
