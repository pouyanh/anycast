package infrastructure

import (
	"bytes"
	"fmt"
)

type LogLevel byte

const (
	FATAL LogLevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)

type LevelledLogger interface {
	Log(level LogLevel, format string, v ...interface{})
}

type MockLevelledLogger struct {
	Level LogLevel

	buffer *bytes.Buffer
}

func (l *MockLevelledLogger) Log(level LogLevel, format string, v ...interface{}) {
	if nil == l.buffer {
		l.buffer = new(bytes.Buffer)
	}

	l.buffer.WriteString(fmt.Sprintf(format, v...))
}

func (l *MockLevelledLogger) Read(p []byte) (int, error) {
	return l.buffer.Read(p)
}

type Logger interface {
	Log(format string, v ...interface{})
}

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
