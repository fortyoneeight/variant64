package entity

import (
	"encoding/json"
	"fmt"
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

func TestRequestRoomAddPlayer(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()

	testcases := []struct {
		name            string
		requests        []RequestRoomAddPlayer
		expectedPlayers []uuid.UUID
	}{
		{
			name: "Add one player.",
			requests: []RequestRoomAddPlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
			},
			expectedPlayers: []uuid.UUID{playerID1},
		},
		{
			name: "Add multiple players.",
			requests: []RequestRoomAddPlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
				{RoomID: uuid.New(), PlayerID: playerID2},
			},
			expectedPlayers: []uuid.UUID{playerID1, playerID2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			room := &Entity[Room]{}
			newRoomRequest := RequestNewRoom{Name: "room1"}
			newRoomRequest.Write(room)

			for _, r := range tc.requests {
				r.Write(room)
			}

			assert.Equal(t, tc.expectedPlayers, room.Data.Players)
		})
	}
}

func TestRequestRoomRemovePlayer(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()

	testcases := []struct {
		name                  string
		addRequests           []RequestRoomAddPlayer
		removeRequests        []RequestRoomRemovePlayer
		expectedPlayersBefore []uuid.UUID
		expectedPlayersAfter  []uuid.UUID
	}{
		{
			name: "Remove one player.",
			addRequests: []RequestRoomAddPlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
			},
			removeRequests: []RequestRoomRemovePlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1},
			expectedPlayersAfter:  []uuid.UUID{},
		},
		{
			name: "Remove one player with remaining.",
			addRequests: []RequestRoomAddPlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
				{RoomID: uuid.New(), PlayerID: playerID2},
			},
			removeRequests: []RequestRoomRemovePlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1, playerID2},
			expectedPlayersAfter:  []uuid.UUID{playerID2},
		},
		{
			name: "Remove multiple players.",
			addRequests: []RequestRoomAddPlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
				{RoomID: uuid.New(), PlayerID: playerID2},
			},
			removeRequests: []RequestRoomRemovePlayer{
				{RoomID: uuid.New(), PlayerID: playerID1},
				{RoomID: uuid.New(), PlayerID: playerID2},
			},
			expectedPlayersBefore: []uuid.UUID{playerID1, playerID2},
			expectedPlayersAfter:  []uuid.UUID{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			room := &Entity[Room]{}
			newRoomRequest := RequestNewRoom{Name: "room1"}
			newRoomRequest.Write(room)

			for _, r := range tc.addRequests {
				r.Write(room)
			}

			assert.Equal(t, tc.expectedPlayersBefore, room.Data.Players)

			for _, r := range tc.removeRequests {
				r.Write(room)
			}

			assert.Equal(t, tc.expectedPlayersAfter, room.Data.Players)
		})
	}
}

func TestSample(t *testing.T) {

	type A struct {
		Val string `json:"val"`
	}

	type B struct {
		A
		Another string `json:"another"`
	}

	b := &B{}
	json.Unmarshal([]byte("{\"val\":\"val\",\"another\":\"another\"}"), b)

	fmt.Println(b)
}