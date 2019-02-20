package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/pouyanh/anycast/lib/actor"
)

type levelled struct {
	logger *logrus.Logger
}

var lvlMap = map[actor.LogLevel]logrus.Level{
	actor.FATAL: logrus.FatalLevel,
	actor.ERROR: logrus.ErrorLevel,
	actor.WARN:  logrus.WarnLevel,
	actor.INFO:  logrus.InfoLevel,
	actor.DEBUG: logrus.DebugLevel,
}

func NewLevelledLogger(lvl actor.LogLevel) (actor.LevelledLogger, error) {
	logger := logrus.New()
	logger.Level = lvlMap[lvl]

	return levelled{logger}, nil
}

func (s levelled) Log(level actor.LogLevel, format string, v ...interface{}) {
	switch level {
	case actor.FATAL:
		s.logger.Fatalf(format, v...)

	case actor.ERROR:
		s.logger.Errorf(format, v...)

	case actor.WARN:
		s.logger.Warnf(format, v...)

	case actor.INFO:
		s.logger.Infof(format, v...)

	case actor.DEBUG:
		s.logger.Debugf(format, v...)

	default:
		s.logger.Printf(format, v...)
	}
}
