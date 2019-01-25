package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/pouyanh/anycast/lib/infrastructure"
)

type levelled struct {
	logger *logrus.Logger
}

var lvlMap = map[infrastructure.LogLevel]logrus.Level {
	infrastructure.FATAL: logrus.FatalLevel,
	infrastructure.ERROR: logrus.ErrorLevel,
	infrastructure.WARN: logrus.WarnLevel,
	infrastructure.INFO: logrus.InfoLevel,
	infrastructure.DEBUG: logrus.DebugLevel,
}

func NewLevelledLogger(lvl infrastructure.LogLevel) (infrastructure.LevelledLogger, error) {
	logger := logrus.New()
	logger.Level = lvlMap[lvl]

	return levelled{logger}, nil
}

func (s levelled) Log(level infrastructure.LogLevel, format string, v ...interface{}) {
	switch level {
	case infrastructure.FATAL:
		s.logger.Fatalf(format, v...)

	case infrastructure.ERROR:
		s.logger.Errorf(format, v...)

	case infrastructure.WARN:
		s.logger.Warnf(format, v...)

	case infrastructure.INFO:
		s.logger.Infof(format, v...)

	case infrastructure.DEBUG:
		s.logger.Debugf(format, v...)

	default:
		s.logger.Printf(format, v...)
	}
}
