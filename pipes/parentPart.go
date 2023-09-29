package pipes

import (
	"bufio"
	"io"
	"os"
	"strings"

	logger "github.com/kalyan3104/dme-logger-go"
)

const logLinesSinkName = "logLinesSink"
const textOutputSinkName = "textOutputSink"

type parentPart struct {
	childName          string
	messenger          *ParentMessenger
	logLinesSink       logger.Logger
	textOutputSink     logger.Logger
	logLineMarshalizer logger.Marshalizer
	loopState          partLoopState

	logsReader    *os.File
	logsWriter    *os.File
	profileReader *os.File
	profileWriter *os.File
}

// NewParentPart creates a new logs receiver part (in the parent process)
func NewParentPart(childName string, logLineMarshalizer logger.Marshalizer) (*parentPart, error) {
	part := &parentPart{
		childName:          childName,
		logLinesSink:       logger.GetOrCreate(logLinesSinkName),
		textOutputSink:     logger.GetOrCreate(textOutputSinkName),
		logLineMarshalizer: logLineMarshalizer,
	}

	err := part.initializePipes()
	if err != nil {
		return nil, err
	}

	part.initializeMessenger()

	return part, nil
}

func (part *parentPart) initializePipes() error {
	var err error

	part.logsReader, part.logsWriter, err = os.Pipe()
	if err != nil {
		return err
	}

	part.profileReader, part.profileWriter, err = os.Pipe()
	if err != nil {
		return err
	}

	return nil
}

func (part *parentPart) initializeMessenger() {
	part.messenger = NewParentMessenger(part.logsReader, part.profileWriter, part.logLineMarshalizer)
}

// GetChildPipes gets the two pipes that should be fed to a child process
func (part *parentPart) GetChildPipes() (*os.File, *os.File) {
	return part.profileReader, part.logsWriter
}

// StartLoop starts the logs reading loop and starts listening for log profile changes (in order to forward them)
func (part *parentPart) StartLoop(childStdout io.Reader, childStderr io.Reader) error {
	if !part.loopState.isInit() {
		return ErrInvalidOperationGivenPartLoopState
	}

	part.loopState.setRunning()

	logger.SubscribeToProfileChange(part)
	part.forwardProfile()
	part.continuouslyRead(childStdout, childStderr)
	return nil
}

// OnProfileChanged is called when a log profile changes
func (part *parentPart) OnProfileChanged() {
	part.forwardProfile()
}

func (part *parentPart) forwardProfile() {
	profile := logger.GetCurrentProfile()
	part.messenger.SendProfile(profile)
}

func (part *parentPart) continuouslyRead(childStdout io.Reader, childStderr io.Reader) {
	go part.continuouslyReadLogLines()
	go part.continuouslyReadStdout(childStdout)
	go part.continuouslyReadStderr(childStderr)
}

func (part *parentPart) continuouslyReadLogLines() {
	for {
		if !part.loopState.isRunning() {
			break
		}

		logLine, err := part.messenger.ReadLogLine()
		if err != nil {
			part.logLinesSink.Error("parentPart.continuouslyReadLogLines()", "err", err)
			break
		}

		part.logLinesSink.Log(logLine)
	}
}

func (part *parentPart) continuouslyReadStdout(stdout io.Reader) {
	stdoutReader := bufio.NewReader(stdout)

	for {
		if !part.loopState.isRunning() {
			break
		}

		textLine, err := stdoutReader.ReadString('\n')
		if err != nil {
			break
		}

		textLine = strings.TrimSpace(textLine)
		part.textOutputSink.Trace(part.childName, "line", textLine)
	}
}

func (part *parentPart) continuouslyReadStderr(stderr io.Reader) {
	stderrReader := bufio.NewReader(stderr)

	for {
		if !part.loopState.isRunning() {
			break
		}

		textLine, err := stderrReader.ReadString('\n')
		if err != nil {
			break
		}

		textLine = strings.TrimSpace(textLine)
		part.textOutputSink.Error(part.childName, "line", textLine)
	}
}

// StopLoop closes all the pipes and stops listening for log profile changes
func (part *parentPart) StopLoop() {
	part.loopState.setStopped()
	logger.UnsubscribeFromProfileChange(part)

	_ = part.logsReader.Close()
	_ = part.logsWriter.Close()
	_ = part.profileReader.Close()
	_ = part.profileWriter.Close()
}
