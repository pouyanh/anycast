package actor

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

type Logger interface {
	Log(format string, v ...interface{})
}
