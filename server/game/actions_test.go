package game

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestNewGameValid(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()
	playerID3 := uuid.New()
	testcases := []struct {
		name                 string
		request              RequestNewGame
		expectedActivePlayer uuid.UUID
		expectedPlayerOrder  []uuid.UUID
	}{
		{
			name: "New game, two players",
			request: RequestNewGame{
				PlayerOrder:     []uuid.UUID{playerID1, playerID2},
				PlayerTimeMilis: 1_000,
			},
			expectedActivePlayer: playerID1,
			expectedPlayerOrder:  []uuid.UUID{playerID2, playerID1},
		},
		{
			name: "New game, three players",
			request: RequestNewGame{
				PlayerOrder:     []uuid.UUID{playerID1, playerID2, playerID3},
				PlayerTimeMilis: 1_000,
			},
			expectedActivePlayer: playerID1,
			expectedPlayerOrder:  []uuid.UUID{playerID2, playerID3, playerID1},
		},
	}

	for _, tc := range testcases {
		t.Run(t.Name(), func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedPlayerOrder, game.playerOrder)
			for _, timer := range game.playerTimers {
				assert.NotNil(t, timer)
			}
		})
	}
}

func TestRequestNewGameInvalid(t *testing.T) {
	playerID1 := uuid.New()
	testcases := []struct {
		name    string
		request RequestNewGame
	}{
		{
			name: "New game without enough players.",
			request: RequestNewGame{
				PlayerOrder:     []uuid.UUID{playerID1},
				PlayerTimeMilis: 1_000,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(t.Name(), func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, game)
			assert.NotNil(t, err)
		})
	}
}

func TestGetGameValid(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()
	game, _ := (&RequestNewGame{
		PlayerOrder:     []uuid.UUID{playerID1, playerID2},
		PlayerTimeMilis: 1_000,
	}).PerformAction()

	testcases := []struct {
		name       string
		request    RequestGetGame
		expectedID uuid.UUID
	}{
		{
			name: "Get player.",
			request: RequestGetGame{
				GameID: game.GetID(),
			},
			expectedID: game.GetID(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedID, game.ID)
		})
	}
}

func TestGetGameInvalid(t *testing.T) {
	testcases := []struct {
		name    string
		request RequestGetGame
	}{
		{
			name: "Get non-existent game.",
			request: RequestGetGame{
				GameID: uuid.New(),
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, game)
			assert.NotNil(t, err)
		})
	}
}

func TestStartGame(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()
	game, _ := (&RequestNewGame{
		PlayerOrder:     []uuid.UUID{playerID1, playerID2},
		PlayerTimeMilis: 1_000,
	}).PerformAction()

	testcases := []struct {
		name       string
		request    RequestStartGame
		expectedID uuid.UUID
	}{
		{
			name: "Start game.",
			request: RequestStartGame{
				GameID: game.GetID(),
			},
			expectedID: game.GetID(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedID, game.ID)
		})
	}
}

func TestStartGameInvalid(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()
	game, _ := (&RequestNewGame{
		PlayerOrder:     []uuid.UUID{playerID1, playerID2},
		PlayerTimeMilis: 1_000,
	}).PerformAction()
	game.start()

	testcases := []struct {
		name    string
		request RequestStartGame
	}{
		{
			name: "Start game.",
			request: RequestStartGame{
				GameID: game.GetID(),
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, game)
			assert.NotNil(t, err)
		})
	}
}

func TestConcedeGame(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()
	playerID3 := uuid.New()
	game, _ := (&RequestNewGame{
		PlayerOrder:     []uuid.UUID{playerID1, playerID2, playerID3},
		PlayerTimeMilis: 1_000,
	}).PerformAction()
	game.start()

	testcases := []struct {
		name            string
		request         RequestConcedeGame
		expectedWinners []uuid.UUID
		expectedLosers  []uuid.UUID
	}{
		{
			name: "Concede game.",
			request: RequestConcedeGame{
				GameID:   game.GetID(),
				PlayerID: playerID1,
			},
			expectedWinners: []uuid.UUID{playerID2, playerID3},
			expectedLosers:  []uuid.UUID{playerID1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedWinners, game.Winners)
			assert.Equal(t, tc.expectedLosers, game.Losers)
			assert.Equal(t, StateFinished, game.State)
		})
	}
}

func TestConcedeGameInvalid(t *testing.T) {
	playerID1 := uuid.New()
	playerID2 := uuid.New()
	playerID3 := uuid.New()
	game1, _ := (&RequestNewGame{
		PlayerOrder:     []uuid.UUID{playerID1, playerID2, playerID3},
		PlayerTimeMilis: 1_000,
	}).PerformAction()
	game2, _ := (&RequestNewGame{
		PlayerOrder:     []uuid.UUID{playerID1, playerID2, playerID3},
		PlayerTimeMilis: 1_000,
	}).PerformAction()
	game2.State = StateFinished

	testcases := []struct {
		name    string
		request RequestConcedeGame
	}{
		{
			name: "Concede game when not started.",
			request: RequestConcedeGame{
				GameID:   game1.GetID(),
				PlayerID: playerID1,
			},
		},
		{
			name: "Concede game when finished.",
			request: RequestConcedeGame{
				GameID:   game2.GetID(),
				PlayerID: playerID1,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := tc.request.PerformAction()

			assert.Nil(t, game)
			assert.NotNil(t, err)
		})
	}
}
