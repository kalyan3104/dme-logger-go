package pipes

import (
	"os"
	"sync"
	"testing"

	logger "github.com/kalyan3104/dme-logger-go"
	"github.com/kalyan3104/dme-logger-go/marshal"
	"github.com/stretchr/testify/require"
)

func TestChildPart_CannotStartLoopTwice(t *testing.T) {
	childPart, err := NewChildPart(os.Stdin, os.Stdout, &marshal.JSONMarshalizer{})
	require.Nil(t, err)

	err = childPart.StartLoop()
	require.Nil(t, err)

	err = childPart.StartLoop()
	require.Equal(t, ErrInvalidOperationGivenPartLoopState, err)
}

func TestChildPart_CannotStartLoopIfStopped(t *testing.T) {
	childPart, err := NewChildPart(os.Stdin, os.Stdout, &marshal.JSONMarshalizer{})
	require.Nil(t, err)

	err = childPart.StartLoop()
	require.Nil(t, err)

	childPart.StopLoop()

	err = childPart.StartLoop()
	require.Equal(t, ErrInvalidOperationGivenPartLoopState, err)
}

func TestChildPart_NoPanicWhenNoParent(t *testing.T) {
	// Bad pipes (no parent)
	profileReader := os.NewFile(4242, "/proc/self/fd/4242")
	logsWriter := os.NewFile(4343, "/proc/self/fd/4343")

	logLineMarshalizer := &marshal.JSONMarshalizer{}
	childLogger := logger.GetOrCreate("child-log")
	childPart, err := NewChildPart(profileReader, logsWriter, logLineMarshalizer)
	require.Nil(t, err)
	err = childPart.StartLoop()
	require.Nil(t, err)

	childLogger.Debug("foo")
	childLogger.Trace("bar")
}

func TestChildPart_ConcurrentWriteLogs(t *testing.T) {
	profileReader := os.NewFile(4242, "/proc/self/fd/4242")
	logsWriter := os.NewFile(4343, "/proc/self/fd/4343")

	childPart, err := NewChildPart(profileReader, logsWriter, &marshal.JSONMarshalizer{})
	require.Nil(t, err)

	err = childPart.StartLoop()
	require.Nil(t, err)

	wg := sync.WaitGroup{}
	wg.Add(2)

	childLogger := logger.GetOrCreate("child-log")

	go func() {
		for i := 0; i < 1000; i++ {
			childLogger.Debug("foo")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			childLogger.Trace("bar")
		}
		wg.Done()
	}()

	wg.Wait()
}
