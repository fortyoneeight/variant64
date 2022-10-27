package entity

import (
	"time"
)

// Timer is a representation of decrementing clock.
type Timer struct {
	timeMilis      time.Duration
	decrementMilis time.Duration

	ticker   *time.Ticker
	running  bool
	doneChan chan bool

	TimerChan chan int64
}

// Start initializes the Timer, it begins paused by default.
func (t *Timer) Start() {
	if !t.running {
		t.setup()
		go t.updateRoutine()
		t.resetTicker()
	}
}

// Stop exits the Timer's decrementing loop.
func (t *Timer) Stop() {
	if t.running {
		t.doneChan <- true
	}
}

// Pause temporarily suspends the Timer from decrementing.
func (t *Timer) Pause() {
	if t.running {
		t.running = false
	}
}

// Unpause unsuspends the Timer.
func (t *Timer) Unpause() {
	if !t.running {
		t.running = true
		t.resetTicker()
	}
}

// setup initializes the Timer.
func (t *Timer) setup() {
	t.doneChan = make(chan bool)
	t.ticker = time.NewTicker(time.Duration(t.decrementMilis))
}

// resetTicker resets the Timer's internal time.Ticker.
func (t *Timer) resetTicker() {
	t.ticker.Reset(time.Duration(t.decrementMilis))
}

// updateRoutine is the loop that handles ticker ticks.
func (t *Timer) updateRoutine() {
	for {
		select {
		case <-t.doneChan:
			return
		case <-t.ticker.C:
			if t.running {
				t.handleTick()
				t.publishTime()
			}
		}
	}
}

// handleTick calculates a Timer's decrement,
// the results are published to the subscribers.
func (t *Timer) handleTick() {
	t.timeMilis -= t.decrementMilis
	if t.timeMilis <= 0 {
		t.timeMilis = 0
		t.running = false
	}
}

// publishTime sends the Timer's value to all subscribers.
func (t *Timer) publishTime() {
	t.TimerChan <- t.timeMilis.Milliseconds()
}

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
