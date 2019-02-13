package logger

import (
	"bytes"
	"fmt"
)

type MockLogger struct {
	buffer *bytes.Buffer
}

func (l *MockLogger) Log(format string, v ...interface{}) {
	if nil == l.buffer {
		l.buffer = new(bytes.Buffer)
	}

	l.buffer.WriteString(fmt.Sprintf(format, v...))
}

func (l *MockLogger) Read(p []byte) (int, error) {
	return l.buffer.Read(p)
}

