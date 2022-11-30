package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestNewPlayer(t *testing.T) {
	testcases := []struct {
		name                string
		request             RequestNewPlayer
		expectedDisplayName string
	}{
		{
			name: "New player.",
			request: RequestNewPlayer{
				DisplayName: "player1",
			},
			expectedDisplayName: "player1",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			entity, err := tc.request.PerformAction()

			assert.Nil(t, err)
			assert.NotNil(t, entity.EntityStore)
			assert.NotNil(t, entity.Data)
			assert.Equal(t, tc.expectedDisplayName, entity.Data.DisplayName)
		})
	}
}

func TestRequestGetPlayer(t *testing.T) {
	player, _ := (&RequestNewPlayer{
		DisplayName: "player1",
	}).PerformAction()

	testcases := []struct {
		name       string
		request    RequestGetPlayer
		expectedID uuid.UUID
	}{
		{
			name: "Get player.",
			request: RequestGetPlayer{
				PlayerID: player.Data.GetID(),
			},
			expectedID: player.Data.GetID(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			entity, err := tc.request.PerformAction()

			assert.Nil(t, err)
			assert.NotNil(t, entity.EntityStore)
			assert.NotNil(t, entity.Data)
			assert.Equal(t, tc.expectedID, entity.Data.ID)
		})
	}
}
