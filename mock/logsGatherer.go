package mock

import (
	"strings"
	"sync"

	logger "github.com/kalyan3104/dme-logger-go"
)

// DummyLogsGatherer -
type DummyLogsGatherer struct {
	lines []logger.LogLineHandler
	text  strings.Builder
	mutex sync.RWMutex
}

// Write -
func (gatherer *DummyLogsGatherer) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Output -
func (gatherer *DummyLogsGatherer) Output(line logger.LogLineHandler) []byte {
	gatherer.mutex.Lock()
	defer gatherer.mutex.Unlock()

	gatherer.lines = append(gatherer.lines, line)
	gatherer.gatherText(line)
	return make([]byte, 0)
}

func (gatherer *DummyLogsGatherer) gatherText(line logger.LogLineHandler) {
	gatherer.text.WriteString(line.GetMessage() + "\n")

	for _, arg := range line.GetArgs() {
		gatherer.text.WriteString(arg + "\n")
	}
}

// GetText -
func (gatherer *DummyLogsGatherer) GetText() string {
	gatherer.mutex.RLock()
	defer gatherer.mutex.RUnlock()

	return gatherer.text.String()
}

// ContainsText -
func (gatherer *DummyLogsGatherer) ContainsText(str string) bool {
	gatherer.mutex.RLock()
	defer gatherer.mutex.RUnlock()

	text := gatherer.text.String()
	return strings.Contains(text, str)
}

// ContainsLogLine -
func (gatherer *DummyLogsGatherer) ContainsLogLine(loggerName string, level logger.LogLevel, message string) bool {
	gatherer.mutex.RLock()
	defer gatherer.mutex.RUnlock()

	for _, line := range gatherer.lines {
		matchedLevel := line.GetLogLevel() == int32(level)
		matchedMessage := line.GetMessage() == message
		matchedLogger := line.GetLoggerName() == loggerName

		if matchedLevel && matchedMessage && matchedLogger {
			return true
		}
	}

	return false
}

// IsInterfaceNil -
func (gatherer *DummyLogsGatherer) IsInterfaceNil() bool {
	return gatherer == nil
}
