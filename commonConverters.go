package logger

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/kalyan3104/dme-logger-go/proto"
)

const bracketsLength = len("[]")
const loggerNameFixedLength = 20
const correlationElementsFixedLength = 14
const messageFixedLength = 40
const ellipsisString = ".."

func displayTime(timestamp int64) string {
	t := time.Unix(0, timestamp)
	return t.Format("2006-01-02 15:04:05.000")
}

func formatMessage(msg string) string {
	return padRight(msg, messageFixedLength)
}

func padRight(str string, maxLength int) string {
	paddingLength := maxLength - len(str)

	if paddingLength > 0 {
		return str + strings.Repeat(" ", paddingLength)
	}

	return str
}

func formatLoggerName(name string) string {
	name = truncatePrefix(name, loggerNameFixedLength-bracketsLength)
	formattedName := fmt.Sprintf("[%s]", name)

	return padRight(formattedName, loggerNameFixedLength)
}

func truncatePrefix(str string, maxLength int) string {
	if len(str) > maxLength {
		startingIndex := len(str) - maxLength + len(ellipsisString)
		return ellipsisString + str[startingIndex:]
	}

	return str
}

func formatCorrelationElements(correlation proto.LogCorrelationMessage) string {
	shard := correlation.GetShard()
	epoch := correlation.GetEpoch()
	round := correlation.GetRound()
	subRound := correlation.GetSubRound()
	formattedElements := fmt.Sprintf("[%s/%d/%d/%s]", shard, epoch, round, subRound)

	return padRight(formattedElements, correlationElementsFixedLength)
}

// ToHexShort generates a short-hand of provided bytes slice showing only the first 3 and the last 3 bytes as hex
// in total, the resulting string is maximum 13 characters long
func ToHexShort(slice []byte) string {
	if len(slice) == 0 {
		return ""
	}
	if len(slice) < 6 {
		return hex.EncodeToString(slice)
	}

	prefix := hex.EncodeToString(slice[:3])
	suffix := hex.EncodeToString(slice[len(slice)-3:])
	return prefix + ellipsisString + suffix
}

// ToHex converts the provided byte slice to its hex represantation
func ToHex(slice []byte) string {
	return hex.EncodeToString(slice)
}
