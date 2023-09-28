package pipes

import "sync/atomic"

// PartLoopStateType represents the state of a part (parent, child) loop
type PartLoopStateType = uint32

const (
	// PartLoopInit signals that loop hasn't been started yet
	PartLoopInit PartLoopStateType = iota
	// PartLoopRunning signals that a loop is running
	PartLoopRunning
	// PartLoopStopped signals that a loop is stopped
	PartLoopStopped
)

type partLoopState struct {
	value PartLoopStateType
}

func (state *partLoopState) setRunning() {
	state.set(PartLoopRunning)
}

func (state *partLoopState) setStopped() {
	state.set(PartLoopStopped)
}

func (state *partLoopState) isInit() bool {
	return state.is(PartLoopInit)
}

func (state *partLoopState) isRunning() bool {
	return state.is(PartLoopRunning)
}

func (state *partLoopState) isStopped() bool {
	return state.is(PartLoopStopped)
}

func (state *partLoopState) set(value PartLoopStateType) {
	atomic.StoreUint32(&state.value, value)
}

func (state *partLoopState) is(value PartLoopStateType) bool {
	return atomic.CompareAndSwapUint32(&state.value, value, value)
}
