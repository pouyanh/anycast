package logger

import (
	"bytes"
	"fmt"

	"github.com/pouyanh/anycast/lib/actor"
)

type MockLevelledLogger struct {
	Level actor.LogLevel

	buffer *bytes.Buffer
}

func (l *MockLevelledLogger) Log(level actor.LogLevel, format string, v ...interface{}) {
	if nil == l.buffer {
		l.buffer = new(bytes.Buffer)
	}

	l.buffer.WriteString(fmt.Sprintf(format, v...))
}

func (l *MockLevelledLogger) Read(p []byte) (int, error) {
	return l.buffer.Read(p)
}
