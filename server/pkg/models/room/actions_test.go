package room

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/models/player"
)

func TestRequestJoinRoom(t *testing.T) {
	newRoomRequest1 := &RequestNewRoom{Name: "room1"}
	room1, err := newRoomRequest1.PerformAction()
	assert.Nil(t, err)

	newRoomRequest2 := &RequestNewRoom{Name: "room2"}
	room2, err := newRoomRequest2.PerformAction()
	assert.Nil(t, err)

	newRoomRequest3 := &RequestNewRoom{Name: "room3"}
	room3, err := newRoomRequest3.PerformAction()
	assert.Nil(t, err)

	player1, err := (&player.RequestNewPlayer{
		DisplayName: "name1",
	}).PerformAction()
	assert.Nil(t, err)

	player2, err := (&player.RequestNewPlayer{
		DisplayName: "name2",
	}).PerformAction()
	assert.Nil(t, err)

	player3, err := (&player.RequestNewPlayer{
		DisplayName: "name3",
	}).PerformAction()
	assert.Nil(t, err)

	playerID1 := player1.ID
	playerID2 := player2.ID
	playerID3 := player3.ID

	testcases := []struct {
		name                  string
		room                  *Room
		requests              []*RequestJoinRoom
		expectedPlayers       map[uuid.UUID]string
		expectedRequestErrors []error
	}{
		{
			name: "Add one player.",
			room: room1,
			requests: []*RequestJoinRoom{
				{RoomID: room1.GetID(), PlayerID: playerID1},
			},
			expectedRequestErrors: []error{nil, nil, nil},
			expectedPlayers: map[uuid.UUID]string{
				playerID1: "name1",
			},
		},
		{
			name: "Add multiple players.",
			room: room2,
			requests: []*RequestJoinRoom{
				{RoomID: room2.GetID(), PlayerID: playerID1},
				{RoomID: room2.GetID(), PlayerID: playerID2},
			},
			expectedRequestErrors: []error{nil, nil, nil},
			expectedPlayers: map[uuid.UUID]string{
				playerID1: "name1",
				playerID2: "name2",
			},
		},
		{
			name: "Add too many players.",
			room: room2,
			requests: []*RequestJoinRoom{
				{RoomID: room3.GetID(), PlayerID: playerID1},
				{RoomID: room3.GetID(), PlayerID: playerID2},
				{RoomID: room3.GetID(), PlayerID: playerID3},
			},
			expectedRequestErrors: []error{nil, nil, errPlayerLimit},
			expectedPlayers: map[uuid.UUID]string{
				playerID1: "name1",
				playerID2: "name2",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for i, r := range tc.requests {
				_, err := r.PerformAction()
				assert.Equal(t, err, tc.expectedRequestErrors[i])
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

	player1, err := (&player.RequestNewPlayer{
		DisplayName: "name1",
	}).PerformAction()
	assert.Nil(t, err)

	player2, err := (&player.RequestNewPlayer{
		DisplayName: "name2",
	}).PerformAction()
	assert.Nil(t, err)

	playerID1 := player1.ID
	playerID2 := player2.ID

	testcases := []struct {
		name                  string
		room                  *Room
		addRequests           []*RequestJoinRoom
		removeRequests        []*RequestLeaveRoom
		expectedPlayersBefore map[uuid.UUID]string
		expectedPlayersAfter  map[uuid.UUID]string
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
			expectedPlayersBefore: map[uuid.UUID]string{
				playerID1: "name1",
			},
			expectedPlayersAfter: map[uuid.UUID]string{},
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
			expectedPlayersBefore: map[uuid.UUID]string{
				playerID1: "name1",
				playerID2: "name2",
			},
			expectedPlayersAfter: map[uuid.UUID]string{
				playerID2: "name2",
			},
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
			expectedPlayersBefore: map[uuid.UUID]string{
				playerID1: "name1",
				playerID2: "name2",
			},
			expectedPlayersAfter: map[uuid.UUID]string{},
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

	player1, err := (&player.RequestNewPlayer{
		DisplayName: "name1",
	}).PerformAction()
	assert.Nil(t, err)

	player2, err := (&player.RequestNewPlayer{
		DisplayName: "name2",
	}).PerformAction()
	assert.Nil(t, err)

	room, err = (&RequestJoinRoom{
		RoomID:   room.GetID(),
		PlayerID: player1.ID,
	}).PerformAction()
	assert.Nil(t, err)

	room, err = (&RequestJoinRoom{
		RoomID:   room.GetID(),
		PlayerID: player2.ID,
	}).PerformAction()
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
				game, err := r.PerformAction()
				assert.Nil(t, err)

				if tc.expectGame {
					assert.NotNil(t, game)
				} else {
					assert.Nil(t, game)
				}
			}
		})
	}
}
