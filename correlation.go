package logger

import (
	"sync"

	"github.com/kalyan3104/dme-logger-go/proto"
)

var globalCorrelation logCorrelation

// logCorrelation holds log correlation elements
type logCorrelation struct {
	mut      sync.RWMutex
	enabled  bool
	shard    string
	epoch    uint32
	round    int64
	subRound string
}

// ToggleCorrelation enables or disables correlation elements for log lines
func ToggleCorrelation(enable bool) {
	globalCorrelation.mut.Lock()
	globalCorrelation.enabled = enable
	globalCorrelation.mut.Unlock()
}

// IsEnabledCorrelation returns whether correlation elements are enabled
func IsEnabledCorrelation() bool {
	globalCorrelation.mut.RLock()
	enabled := globalCorrelation.enabled
	globalCorrelation.mut.RUnlock()

	return enabled
}

// SetCorrelationShard sets the current shard ID as a log correlation element
func SetCorrelationShard(shardID string) {
	globalCorrelation.mut.Lock()
	globalCorrelation.shard = shardID
	globalCorrelation.mut.Unlock()
}

// SetCorrelationEpoch sets the current epoch as a log correlation element
func SetCorrelationEpoch(epoch uint32) {
	globalCorrelation.mut.Lock()
	globalCorrelation.epoch = epoch
	globalCorrelation.mut.Unlock()
}

// SetCorrelationRound sets the current round as a log correlation element
func SetCorrelationRound(round int64) {
	globalCorrelation.mut.Lock()
	globalCorrelation.round = round
	globalCorrelation.mut.Unlock()
}

// SetCorrelationSubround sets the current sub-round as a log correlation element
func SetCorrelationSubround(subRound string) {
	globalCorrelation.mut.Lock()
	globalCorrelation.subRound = subRound
	globalCorrelation.mut.Unlock()
}

// GetCorrelation gets global correlation elements
func GetCorrelation() proto.LogCorrelationMessage {
	globalCorrelation.mut.RLock()
	lcm := proto.LogCorrelationMessage{
		Shard:    globalCorrelation.shard,
		Epoch:    globalCorrelation.epoch,
		Round:    globalCorrelation.round,
		SubRound: globalCorrelation.subRound,
	}
	globalCorrelation.mut.RUnlock()

	return lcm
}
