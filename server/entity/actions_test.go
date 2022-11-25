package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestRoomAddPlayer(t *testing.T) {
	requestHandler := RequestHandler{}

	newRoomRequest1 := &RequestNewRoom{Name: "room1"}
	room1, err := requestHandler.HandleNewRoom(newRoomRequest1)
	assert.Nil(t, err)

	newRoomRequest2 := &RequestNewRoom{Name: "room2"}
	room2, err := requestHandler.HandleNewRoom(newRoomRequest2)
	assert.Nil(t, err)

	playerID1 := uuid.New()
	playerID2 := uuid.New()

	testcases := []struct {
		name            string
		room            *Entity[Room]
		requests        []*RequestRoomAddPlayer
		expectedPlayers []uuid.UUID
	}{
		{
			name: "Add one player.",
			room: room1,
			requests: []*RequestRoomAddPlayer{
				{RoomID: room1.Data.GetID(), PlayerID: playerID1},
			},
			expectedPlayers: []uuid.UUID{playerID1},
		},
		{
			name: "Add multiple players.",
			room: room2,
			requests: []*RequestRoomAddPlayer{
				{RoomID: room2.Data.GetID(), PlayerID: playerID1},
				{RoomID: room2.Data.GetID(), PlayerID: playerID2},
			},
			expectedPlayers: []uuid.UUID{playerID1, playerID2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, r := range tc.requests {
				room, err := requestHandler.HandleRoomAddPlayer(r)
				assert.Nil(t, err)
				tc.room = room
			}

			assert.Equal(t, tc.expectedPlayers, tc.room.Data.Players)
		})
	}
}

func TestRequestRoomRemovePlayer(t *testing.T) {
	requestHandler := RequestHandler{}

	newRoomRequest1 := &RequestNewRoom{Name: "room1"}
	room1, err := requestHandler.HandleNewRoom(newRoomRequest1)
	assert.Nil(t, err)

	newRoomRequest2 := &RequestNewRoom{Name: "room2"}
	room2, err := requestHandler.HandleNewRoom(newRoomRequest2)
	assert.Nil(t, err)

	newRoomRequest3 := &RequestNewRoom{Name: "room3"}
	room3, err := requestHandler.HandleNewRoom(newRoomRequest3)
	assert.Nil(t, err)

	playerID1 := uuid.New()
	playerID2 := uuid.New()

	testcases := []struct {
		name                  string
		room                  *Entity[Room]
		addRequests           []*RequestRoomAddPlayer
		removeRequests        []*RequestRoomRemovePlayer
		expectedPlayersBefore []uuid.UUID
		expectedPlayersAfter  []uuid.UUID
	}{
		{
			name: "Remove one player.",
			room: room1,
			addRequests: []*RequestRoomAddPlayer{
				{RoomID: room1.Data.GetID(), PlayerID: playerID1},
			},
			removeRequests: []*RequestRoomRemovePlayer{
				{RoomID: room1.Data.GetID(), PlayerID: playerID1},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1},
			expectedPlayersAfter:  []uuid.UUID{},
		},
		{
			name: "Remove one player with remaining.",
			room: room2,
			addRequests: []*RequestRoomAddPlayer{
				{RoomID: room2.Data.GetID(), PlayerID: playerID1},
				{RoomID: room2.Data.GetID(), PlayerID: playerID2},
			},
			removeRequests: []*RequestRoomRemovePlayer{
				{RoomID: room2.Data.GetID(), PlayerID: playerID1},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1, playerID2},
			expectedPlayersAfter:  []uuid.UUID{playerID2},
		},
		{
			name: "Remove multiple players.",
			room: room3,
			addRequests: []*RequestRoomAddPlayer{
				{RoomID: room3.Data.GetID(), PlayerID: playerID1},
				{RoomID: room3.Data.GetID(), PlayerID: playerID2},
			},
			removeRequests: []*RequestRoomRemovePlayer{
				{RoomID: room3.Data.GetID(), PlayerID: playerID1},
				{RoomID: room3.Data.GetID(), PlayerID: playerID2},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1, playerID2},
			expectedPlayersAfter:  []uuid.UUID{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, r := range tc.addRequests {
				room, err := requestHandler.HandleRoomAddPlayer(r)
				assert.Nil(t, err)
				tc.room = room
			}
			assert.Equal(t, tc.expectedPlayersBefore, tc.room.Data.Players)

			for _, r := range tc.removeRequests {
				room, err := requestHandler.HandleRoomRemovePlayer(r)
				assert.Nil(t, err)
				tc.room = room
			}
			assert.Equal(t, tc.expectedPlayersAfter, tc.room.Data.Players)
		})
	}
}

func TestRequestRoomStartGame(t *testing.T) {
	requestHandler := RequestHandler{}
	room, err := requestHandler.HandleNewRoom(&RequestNewRoom{Name: "name"})
	assert.Nil(t, err)

	room, err = requestHandler.HandleRoomAddPlayer(&RequestRoomAddPlayer{
		RoomID:   room.Data.GetID(),
		PlayerID: uuid.New(),
	})
	assert.Nil(t, err)

	room, err = requestHandler.HandleRoomAddPlayer(&RequestRoomAddPlayer{
		RoomID:   room.Data.GetID(),
		PlayerID: uuid.New(),
	})
	assert.Nil(t, err)

	testcases := []struct {
		name           string
		requests       []*RequestRoomStartGame
		expectGame     bool
		expectedGameID *uuid.UUID
	}{
		{
			name:       "No GameID.",
			requests:   []*RequestRoomStartGame{},
			expectGame: false,
		},
		{
			name: "Add GameID.",
			requests: []*RequestRoomStartGame{
				{RoomID: room.Data.GetID(), PlayerTimeMilis: 1_000},
			},
			expectGame: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, r := range tc.requests {
				room, err := requestHandler.HandleRoomStartGame(r)
				assert.Nil(t, err)

				if tc.expectGame {
					assert.NotNil(t, room.Data.GameID)
				} else {
					assert.Nil(t, room.Data.GameID)
				}
			}
		})
	}
}
