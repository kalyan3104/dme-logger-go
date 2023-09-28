package logger

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatMessage_ShouldOutputFixedOrLargerStringThanMsgFixedLength(t *testing.T) {
	t.Parallel()

	testData := make(map[string]int)
	testData[""] = 40
	testData["small string"] = 40
	largeString := "large string: " + strings.Repeat("A", messageFixedLength)
	testData[largeString] = len(largeString)

	for k, v := range testData {
		result := formatMessage(k)

		assert.Equal(t, v, len(result))
	}
}

func TestToHexShort_EmptySliceShouldReturnEmptyString(t *testing.T) {
	t.Parallel()

	hash := []byte("")
	res := ToHexShort(hash)

	assert.Equal(t, 0, len(res))
}

func TestToHexShort_SliceLengthSmallShouldNotChange(t *testing.T) {
	t.Parallel()

	input := []byte("short")
	hexHash := hex.EncodeToString(input)

	res := ToHexShort(input)

	assert.Equal(t, hexHash, res)
}

func TestToHexShort_SliceLengthBigShouldTrim(t *testing.T) {
	t.Parallel()

	input := []byte("long enough input so it should be trimmed")
	hexHash := hex.EncodeToString(input)

	res := ToHexShort(input)

	assert.NotEqual(t, len(hexHash), len(res))
	assert.True(t, strings.Contains(res, ellipsisString))
}
