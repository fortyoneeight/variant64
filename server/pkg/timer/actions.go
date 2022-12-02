package timer

import "time"

// RequestNewTimer is used to create a new Timer.
type RequestNewTimer struct {
	StartingTimeMilis int64
	DecrementMilis    int64
}

// NewTimer returns a new Timer created via the provided RequestNewTimer.
func NewTimer(r RequestNewTimer) *Timer {
	return &Timer{
		timeMilis:      time.Duration(r.StartingTimeMilis) * time.Millisecond,
		decrementMilis: time.Duration(r.DecrementMilis) * time.Millisecond,
		running:        false,
		TimerChan:      make(chan int64),
	}
}
