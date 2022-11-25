package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestNewRoom(t *testing.T) {
	testcases := []struct {
		name         string
		request      RequestNewRoom
		expectedName string
	}{
		{
			name: "New empty room.",
			request: RequestNewRoom{
				Name: "room1",
			},
			expectedName: "room1",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			entity := &Entity[Room]{}
			tc.request.Write(entity)

			assert.NotNil(t, entity.EntityStore)
			assert.NotNil(t, entity.Data)
			assert.Equal(t, tc.expectedName, entity.Data.Name)
			assert.Empty(t, entity.Data.Players)
		})
	}
}

func TestRequestGetRoom(t *testing.T) {
	roomID := uuid.New()

	testcases := []struct {
		name       string
		request    RequestGetRoom
		expectedID uuid.UUID
	}{
		{
			name: "New empty room.",
			request: RequestGetRoom{
				ID: roomID,
			},
			expectedID: roomID,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			entity := &Entity[Room]{}
			tc.request.Read(entity)

			assert.NotNil(t, entity.EntityStore)
			assert.NotNil(t, entity.Data)
			assert.Equal(t, tc.expectedID, entity.Data.ID)
		})
	}
}
