package integration

import (
	"fmt"

	"github.com/google/uuid"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/models"
	"github.com/variant64/server/pkg/models/room"
)

func TestIntegrationJoinRoomSubscription(t *testing.T) {
	newRoomRequest1 := &room.RequestNewRoom{Name: "room1"}
	room1, err := newRoomRequest1.PerformAction()
	assert.Nil(t, err)

	playerID1 := uuid.New()

	testcases := []struct {
		name             string
		room             *room.Room
		requests         []*room.RequestJoinRoom
		expectedMessages int
	}{
		{
			name: "Add one player.",
			room: room1,
			requests: []*room.RequestJoinRoom{
				{RoomID: room1.GetID(), PlayerID: playerID1},
			},
			expectedMessages: 3,
		},
	}

	for i, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			newRoom, err := (&room.RequestNewRoom{Name: fmt.Sprint("room_", i)}).PerformAction()
			if err != nil {
				t.Error(err)
			}

			mockWriter := models.NewMockEventWriter()
			err = (&room.CommandRoomSubscribe{RoomID: newRoom.ID, EventWriter: mockWriter}).PerformAction()
			if err != nil {
				t.Error(err)
			}

			newRoom, err = (&room.RequestJoinRoom{RoomID: newRoom.ID, PlayerID: uuid.New()}).PerformAction()
			if err != nil {
				t.Error(err)
			}

			(&room.RequestLeaveRoom{RoomID: newRoom.ID, PlayerID: newRoom.Players[0]}).PerformAction()
			assert.Nil(t, err)
			tc.room = newRoom

			assert.Equal(t, len(mockWriter.SentMessages), 3)
		})
	}
}
