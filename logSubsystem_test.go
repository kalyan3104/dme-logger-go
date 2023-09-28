package logger_test

import (
	"testing"

	logger "github.com/kalyan3104/dme-logger-go"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel_WrongStringParameterShouldErr(t *testing.T) {
	err := logger.SetLogLevel("wrong string")

	assert.Equal(t, logger.ErrInvalidLogLevelPattern, err)
}

func TestSetLogLevel_WrongLogLevelShouldErr(t *testing.T) {
	err := logger.SetLogLevel("*:WRONG LEVEL")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unknown log level")
}

func TestSetLogLevel_NonWildcardPatternShouldNotError(t *testing.T) {
	log1 := logger.GetOrCreate("pattern1")
	log2 := logger.GetOrCreate("pattern2")

	err := logger.SetLogLevel("pattern1:DEbUG")

	assert.Nil(t, err)
	assert.Equal(t, logger.LogDebug, log1.LogLevel())
	assert.Equal(t, logger.LogInfo, log2.LogLevel())

	err = logger.SetLogLevel("pattern:tRace")

	assert.Nil(t, err)
	assert.Equal(t, logger.LogTrace, log1.LogLevel())
	assert.Equal(t, logger.LogTrace, log2.LogLevel())

	// rollback to the default value
	_ = logger.SetLogLevel("*:INFO")
}

func TestSetLogLevel_WildcardPatternShouldWork(t *testing.T) {
	log1 := logger.GetOrCreate("1")
	log2 := logger.GetOrCreate("2")

	err := logger.SetLogLevel("*:DEBuG")

	assert.Nil(t, err)
	assert.Equal(t, logger.LogDebug, log1.LogLevel())
	assert.Equal(t, logger.LogDebug, log2.LogLevel())
	assert.Equal(t, logger.LogDebug, *logger.DefaultLogLevel)

	// rollback to the default value
	_ = logger.SetLogLevel("*:INFO")
}

func TestSetLogLevel_MultipleVariantsShouldWork(t *testing.T) {
	log1 := logger.GetOrCreate("1")
	log2 := logger.GetOrCreate("2")

	err := logger.SetLogLevel("*:DEBuG,1:INFO,*:TRacE,2:ERROR")

	assert.Nil(t, err)
	assert.Equal(t, logger.LogTrace, log1.LogLevel())
	assert.Equal(t, logger.LogError, log2.LogLevel())
	assert.Equal(t, logger.LogTrace, *logger.DefaultLogLevel)

	// rollback to the default value
	_ = logger.SetLogLevel("*:INFO")
}

func TestGetLoggerLogLevel(t *testing.T) {
	_ = logger.GetOrCreate("1")
	_ = logger.GetOrCreate("2")

	err := logger.SetLogLevel("1:INFO,2:TRACE")

	assert.Nil(t, err)
	assert.Equal(t, logger.LogInfo, logger.GetLoggerLogLevel("1"))
	assert.Equal(t, logger.LogTrace, logger.GetLoggerLogLevel("2"))
	assert.Equal(t, logger.LogNone, logger.GetLoggerLogLevel("42"))

	// rollback to the default value
	_ = logger.SetLogLevel("*:INFO")
}
