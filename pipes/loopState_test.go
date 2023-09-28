package pipes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPartLoopState(t *testing.T) {
	state := partLoopState{}
	require.True(t, state.isInit())
	require.False(t, state.isRunning())
	require.False(t, state.isStopped())

	state.setRunning()
	require.False(t, state.isInit())
	require.True(t, state.isRunning())
	require.False(t, state.isStopped())

	state.setStopped()
	require.False(t, state.isInit())
	require.False(t, state.isRunning())
	require.True(t, state.isStopped())
}
