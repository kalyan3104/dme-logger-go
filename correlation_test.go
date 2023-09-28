package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCorrelation_Toggle(t *testing.T) {
	ToggleCorrelation(true)
	require.True(t, IsEnabledCorrelation())

	ToggleCorrelation(false)
	require.False(t, IsEnabledCorrelation())
}

func TestCorrelation_SettingElements(t *testing.T) {
	shard := "myshard"
	epoch := uint32(42)
	round := int64(420)
	subRound := "foo"

	// Now with the global setters
	SetCorrelationShard(shard)
	SetCorrelationEpoch(epoch)
	SetCorrelationRound(round)
	SetCorrelationSubround(subRound)

	lcm := GetCorrelation()

	require.Equal(t, shard, lcm.Shard)
	require.Equal(t, epoch, lcm.Epoch)
	require.Equal(t, round, lcm.Round)
	require.Equal(t, subRound, lcm.SubRound)
}
