package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/pouyanh/anycast/lib/port"
)

type levelled struct {
	logger *logrus.Logger
}

var lvlMap = map[port.LogLevel]logrus.Level{
	port.FATAL: logrus.FatalLevel,
	port.ERROR: logrus.ErrorLevel,
	port.WARN:  logrus.WarnLevel,
	port.INFO:  logrus.InfoLevel,
	port.DEBUG: logrus.DebugLevel,
}

func NewLevelledLogger(lvl port.LogLevel) (port.LevelledLogger, error) {
	logger := logrus.New()
	logger.Level = lvlMap[lvl]

	return levelled{logger}, nil
}

func (s levelled) Log(level port.LogLevel, format string, v ...interface{}) {
	switch level {
	case port.FATAL:
		s.logger.Fatalf(format, v...)

	case port.ERROR:
		s.logger.Errorf(format, v...)

	case port.WARN:
		s.logger.Warnf(format, v...)

	case port.INFO:
		s.logger.Infof(format, v...)

	case port.DEBUG:
		s.logger.Debugf(format, v...)

	default:
		s.logger.Printf(format, v...)
	}
}
