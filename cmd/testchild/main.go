package main

import (
	"fmt"
	goLog "log"
	"os"

	logger "github.com/kalyan3104/dme-logger-go"
	"github.com/kalyan3104/dme-logger-go/marshal"
	"github.com/kalyan3104/dme-logger-go/mock"
	"github.com/kalyan3104/dme-logger-go/pipes"
)

const (
	fileDescriptorProfileReader = 3
	fileDescriptorLogsWriter    = 4
)

func main() {
	profileReader := getPipeFile(fileDescriptorProfileReader)
	if profileReader == nil {
		goLog.Fatal("Cannot get pipe file: [profileReader]")
	}

	logsWriter := getPipeFile(fileDescriptorLogsWriter)
	if logsWriter == nil {
		goLog.Fatal("Cannot get pipe file: [logsWriter]")
	}

	part, err := pipes.NewChildPart(profileReader, logsWriter, &marshal.JSONMarshalizer{})
	if err != nil {
		goLog.Fatal("Can't create part")
	}

	err = part.StartLoop()
	if err != nil {
		goLog.Fatal("Ended loop")
	}

	fooLog := logger.GetOrCreate("foo")
	barLog := logger.GetOrCreate("bar")

	fooLog.Info("foo-info")
	barLog.Info("bar-info")

	fooLog.Trace("foo-trace-no")
	barLog.Trace("bar-trace-no")

	go func() {
		fooLog.Info("foo-in-go")
		barLog.Info("bar-in-go")
	}()

	mock.SendDummySignal("done-step-1")
	mock.WaitUntilLogLevelPattern("*:TRACE")

	fooLog.Trace("foo-trace-yes")
	barLog.Trace("bar-trace-yes")

	fmt.Println("Here's some stdout")
	fmt.Fprintln(os.Stderr, "Here's some stderr")

	mock.SendDummySignal("done-step-2")
}

func getPipeFile(fileDescriptor uintptr) *os.File {
	file := os.NewFile(fileDescriptor, fmt.Sprintf("/proc/self/fd/%d", fileDescriptor))
	return file
}
