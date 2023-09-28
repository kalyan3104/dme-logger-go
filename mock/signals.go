package mock

import (
	"io/ioutil"
	"os"
	"path"
	"time"

	logger "github.com/kalyan3104/dme-logger-go"
)

const dummySignalsFolder = "dummysignals"

// SendDummySignal sends a dummy signal (creates a file)
func SendDummySignal(signal string) {
	os.MkdirAll(dummySignalsFolder, os.ModePerm)

	err := ioutil.WriteFile(getSignalFile(signal), []byte("foobar"), 0755)
	if err != nil {
		panic("Could not send dummy signal")
	}
}

// WaitForDummySignal waits for a signal (a dummy file) to appear
func WaitForDummySignal(signal string) {
	for {
		if _, err := os.Stat(getSignalFile(signal)); err == nil {
			break
		}

		waitABit()
	}
}

// ClearAllDummySignals clears all dummy signals
func ClearAllDummySignals() {
	os.RemoveAll(dummySignalsFolder)
	waitABit()
}

// WaitUntilLogLevelPattern waits until a log level matches the value
func WaitUntilLogLevelPattern(value string) {
	for {
		currentPattern := logger.GetLogLevelPattern()
		if currentPattern == value {
			break
		}

		waitABit()
	}
}

func getSignalFile(signal string) string {
	return path.Join(dummySignalsFolder, signal)
}

func waitABit() {
	time.Sleep(100 * time.Millisecond)
}
