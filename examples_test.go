package logger_test

import (
	"crypto/rand"
	"testing"

	logger "github.com/kalyan3104/dme-logger-go"
)

func TestLogger_ExampleCreateLoggerAndOutputSimpleMessages(t *testing.T) {
	//the following instruction might be done inside a var declaration, once on each package
	// or in the init func of the package
	log := logger.GetOrCreate("test_logger1")
	//manual set of the log lev is required here for demonstration purposes
	log.SetLevel(logger.LogTrace)

	log.Trace("a trace message")
	log.Debug("a debug message")
	log.Info("an information message")
	log.Warn("a warning message")
	log.Error("an error message")
}

func TestLogger_ExampleMessagesWithArguments(t *testing.T) {
	log := logger.GetOrCreate("test_logger2")
	log.SetLevel(logger.LogInfo)

	log.Info("message1", "an-int", 45, "a-string", "string")
	log.Info("message2", "a-map", map[string]int{"key1": 0, "key2": 1})
	log.Info("message3", "a-slice", []int{1, 2, 3, 4, 5})
	log.Info("message4", "nil", nil)
	hash := generateHash()
	log.Info("message5", "hash", hash)
}

func generateHash() []byte {
	buff := make([]byte, 32)
	_, _ = rand.Reader.Read(buff)
	return buff
}
