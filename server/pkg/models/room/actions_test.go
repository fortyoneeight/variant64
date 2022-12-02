package room

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestRoomAddPlayer(t *testing.T) {
	newRoomRequest1 := &RequestNewRoom{Name: "room1"}
	room1, err := newRoomRequest1.PerformAction()
	assert.Nil(t, err)

	newRoomRequest2 := &RequestNewRoom{Name: "room2"}
	room2, err := newRoomRequest2.PerformAction()
	assert.Nil(t, err)

	playerID1 := uuid.New()
	playerID2 := uuid.New()

	testcases := []struct {
		name            string
		room            *Room
		requests        []*RequestJoinRoom
		expectedPlayers []uuid.UUID
	}{
		{
			name: "Add one player.",
			room: room1,
			requests: []*RequestJoinRoom{
				{RoomID: room1.GetID(), PlayerID: playerID1},
			},
			expectedPlayers: []uuid.UUID{playerID1},
		},
		{
			name: "Add multiple players.",
			room: room2,
			requests: []*RequestJoinRoom{
				{RoomID: room2.GetID(), PlayerID: playerID1},
				{RoomID: room2.GetID(), PlayerID: playerID2},
			},
			expectedPlayers: []uuid.UUID{playerID1, playerID2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, r := range tc.requests {
				room, err := r.PerformAction()
				assert.Nil(t, err)
				tc.room = room
			}

			assert.Equal(t, tc.expectedPlayers, tc.room.Players)
		})
	}
}

func TestRequestRoomRemovePlayer(t *testing.T) {
	newRoomRequest1 := &RequestNewRoom{Name: "room1"}
	room1, err := newRoomRequest1.PerformAction()
	assert.Nil(t, err)

	newRoomRequest2 := &RequestNewRoom{Name: "room2"}
	room2, err := newRoomRequest2.PerformAction()
	assert.Nil(t, err)

	newRoomRequest3 := &RequestNewRoom{Name: "room3"}
	room3, err := newRoomRequest3.PerformAction()
	assert.Nil(t, err)

	playerID1 := uuid.New()
	playerID2 := uuid.New()

	testcases := []struct {
		name                  string
		room                  *Room
		addRequests           []*RequestJoinRoom
		removeRequests        []*RequestLeaveRoom
		expectedPlayersBefore []uuid.UUID
		expectedPlayersAfter  []uuid.UUID
	}{
		{
			name: "Remove one player.",
			room: room1,
			addRequests: []*RequestJoinRoom{
				{RoomID: room1.GetID(), PlayerID: playerID1},
			},
			removeRequests: []*RequestLeaveRoom{
				{RoomID: room1.GetID(), PlayerID: playerID1},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1},
			expectedPlayersAfter:  []uuid.UUID{},
		},
		{
			name: "Remove one player with remaining.",
			room: room2,
			addRequests: []*RequestJoinRoom{
				{RoomID: room2.GetID(), PlayerID: playerID1},
				{RoomID: room2.GetID(), PlayerID: playerID2},
			},
			removeRequests: []*RequestLeaveRoom{
				{RoomID: room2.GetID(), PlayerID: playerID1},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1, playerID2},
			expectedPlayersAfter:  []uuid.UUID{playerID2},
		},
		{
			name: "Remove multiple players.",
			room: room3,
			addRequests: []*RequestJoinRoom{
				{RoomID: room3.GetID(), PlayerID: playerID1},
				{RoomID: room3.GetID(), PlayerID: playerID2},
			},
			removeRequests: []*RequestLeaveRoom{
				{RoomID: room3.GetID(), PlayerID: playerID1},
				{RoomID: room3.GetID(), PlayerID: playerID2},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1, playerID2},
			expectedPlayersAfter:  []uuid.UUID{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, r := range tc.addRequests {
				room, err := r.PerformAction()
				assert.Nil(t, err)
				tc.room = room
			}
			assert.Equal(t, tc.expectedPlayersBefore, tc.room.Players)

			for _, r := range tc.removeRequests {
				room, err := r.PerformAction()
				assert.Nil(t, err)
				tc.room = room
			}
			assert.Equal(t, tc.expectedPlayersAfter, tc.room.Players)
		})
	}
}

func TestRequestRoomStartGame(t *testing.T) {
	room, err := (&RequestNewRoom{Name: "name"}).PerformAction()
	assert.Nil(t, err)

	request := &RequestJoinRoom{
		RoomID:   room.GetID(),
		PlayerID: uuid.New(),
	}

	room, err = request.PerformAction()

	assert.Nil(t, err)

	request = &RequestJoinRoom{
		RoomID:   room.GetID(),
		PlayerID: uuid.New(),
	}

	room, err = request.PerformAction()
	assert.Nil(t, err)

	testcases := []struct {
		name           string
		requests       []*RequestStartGame
		expectGame     bool
		expectedGameID *uuid.UUID
	}{
		{
			name:       "No GameID.",
			requests:   []*RequestStartGame{},
			expectGame: false,
		},
		{
			name: "Add GameID.",
			requests: []*RequestStartGame{
				{RoomID: room.GetID(), PlayerTimeMilis: 1_000},
			},
			expectGame: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, r := range tc.requests {
				room, err := r.PerformAction()
				assert.Nil(t, err)

				if tc.expectGame {
					assert.NotNil(t, room.GameID)
				} else {
					assert.Nil(t, room.GameID)
				}
			}
		})
	}
}
