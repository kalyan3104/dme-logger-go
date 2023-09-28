package logger

import "io"

var DefaultLogLevel = &defaultLogLevel

func (los *logOutputSubject) Observers() ([]io.Writer, []Formatter) {
	los.mutObservers.RLock()
	defer los.mutObservers.RUnlock()

	return los.writers, los.formatters
}

func (l *logger) LogLevel() LogLevel {
	return l.logLevel
}
