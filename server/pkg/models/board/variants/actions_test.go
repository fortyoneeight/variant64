package variants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestNewClassicBoard(t *testing.T) {
	classicBoard, err := (&RequestNewClassicBoard{}).PerformAction()
	assert.Nil(t, err)
	assert.NotNil(t, classicBoard)
}
