package timer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimerHandleTick(t *testing.T) {
	testcases := []struct {
		name                  string
		request               RequestNewTimer
		numTicks              int
		expectedTimeSnapshots []int64
		expectedEnd           time.Duration
		shouldBeRunning       bool
	}{
		{
			name: "Timer decrements with remaining.",
			request: RequestNewTimer{
				StartingTimeMilis: 1000,
				DecrementMilis:    500,
			},
			numTicks:              1,
			expectedTimeSnapshots: []int64{500},
			expectedEnd:           time.Duration(500) * time.Millisecond,
			shouldBeRunning:       true,
		},
		{
			name: "Timer decrements multiple with remaining.",
			request: RequestNewTimer{
				StartingTimeMilis: 1000,
				DecrementMilis:    200,
			},
			numTicks:              2,
			expectedTimeSnapshots: []int64{800, 600},
			expectedEnd:           time.Duration(600) * time.Millisecond,
			shouldBeRunning:       true,
		},
		{
			name: "Timer decrements without remaining.",
			request: RequestNewTimer{
				StartingTimeMilis: 1000,
				DecrementMilis:    1000,
			},
			numTicks:              1,
			expectedTimeSnapshots: []int64{0},
			expectedEnd:           time.Duration(0),
			shouldBeRunning:       false,
		},
		{
			name: "Timer decrements multiple without remaining.",
			request: RequestNewTimer{
				StartingTimeMilis: 1000,
				DecrementMilis:    500,
			},
			numTicks:              2,
			expectedTimeSnapshots: []int64{500, 0},
			expectedEnd:           time.Duration(0),
			shouldBeRunning:       false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			timer := NewTimer(tc.request)
			timer.running = true

			timeSnapshot := make([]int64, 0)
			for i := 0; i < tc.numTicks; i++ {
				timer.handleTick()
				timeSnapshot = append(timeSnapshot, timer.timeMilis.Milliseconds())
			}

			assert.Equal(t, timer.running, tc.shouldBeRunning)
			assert.Equal(t, timer.timeMilis, tc.expectedEnd)
			assert.Equal(t, timeSnapshot, tc.expectedTimeSnapshots)
		})
	}
}
