package pipes

import (
	"os"
	"time"

	logger "github.com/kalyan3104/dme-logger-go"
)

// ParentMessenger is the messenger on parent's part of the pipe
type ParentMessenger struct {
	Messenger
	logLineMarshalizer logger.Marshalizer
}

// NewParentMessenger creates a new messenger
func NewParentMessenger(logsReader *os.File, profileWriter *os.File, logLineMarshalizer logger.Marshalizer) *ParentMessenger {
	return &ParentMessenger{
		Messenger:          *NewMessenger(logsReader, profileWriter),
		logLineMarshalizer: logLineMarshalizer,
	}
}

// ReadLogLine reads a log line
func (messenger *ParentMessenger) ReadLogLine() (*logger.LogLine, error) {
	buffer, err := messenger.ReadMessage()
	if err != nil {
		return nil, err
	}

	wrapper := &logger.LogLineWrapper{}
	err = messenger.logLineMarshalizer.Unmarshal(wrapper, buffer)
	if err != nil {
		return nil, err
	}

	logLine := messenger.recoverLogLine(wrapper)
	return logLine, nil
}

func (messenger *ParentMessenger) recoverLogLine(wrapper *logger.LogLineWrapper) *logger.LogLine {
	logLine := &logger.LogLine{
		LoggerName:  wrapper.LoggerName,
		Correlation: wrapper.Correlation,
		Message:     wrapper.Message,
		LogLevel:    logger.LogLevel(wrapper.LogLevel),
		Args:        make([]interface{}, len(wrapper.Args)),
		Timestamp:   time.Unix(0, wrapper.Timestamp),
	}

	for i, str := range wrapper.Args {
		logLine.Args[i] = str
	}

	return logLine
}

// SendProfile sends a profile
func (messenger *ParentMessenger) SendProfile(profile logger.Profile) error {
	buffer, err := profile.Marshal()
	if err != nil {
		return err
	}

	_, err = messenger.SendMessage(buffer)
	if err != nil {
		return err
	}

	return nil
}
