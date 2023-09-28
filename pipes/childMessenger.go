package pipes

import (
	"os"

	logger "github.com/kalyan3104/dme-logger-go"
)

// ChildMessenger is the messenger on child's part of the pipe
type ChildMessenger struct {
	Messenger
}

// NewChildMessenger creates a new messenger
func NewChildMessenger(profileReader *os.File, logsWriter *os.File) *ChildMessenger {
	return &ChildMessenger{
		Messenger: *NewMessenger(profileReader, logsWriter),
	}
}

// ReadProfile reads an incoming profile
func (messenger *ChildMessenger) ReadProfile() (logger.Profile, error) {
	buffer, err := messenger.ReadMessage()
	if err != nil {
		return logger.Profile{}, err
	}

	return logger.UnmarshalProfile(buffer)
}

// SendLogLine sends a log line
func (messenger *ChildMessenger) SendLogLine(logLineMarshalized []byte) (int, error) {
	return messenger.SendMessage(logLineMarshalized)
}
