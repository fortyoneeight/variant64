package game

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/errortypes"
	"github.com/variant64/server/timer"
)

// RequestNewGame is a used to create a new Game.
type RequestNewGame struct {
	PlayerOrder     []uuid.UUID `json:"player_order" swaggerignore:"true"`
	PlayerTimeMilis int64       `json:"player_time_ms"`
}

// PerformAction creates a new Game.
func (r *RequestNewGame) PerformAction() (*Game, errortypes.TypedError) {
	if len(r.PlayerOrder) < 2 {
		return nil, errInvalidPlayersNumber{number: len(r.PlayerOrder)}
	}

	game := &Game{
		ID:           uuid.New(),
		ActivePlayer: r.PlayerOrder[0],
		playerOrder:  append(r.PlayerOrder[1:], r.PlayerOrder[0]),
		playerTimers: make(map[uuid.UUID]*timer.Timer),
		Winners:      []uuid.UUID{},
		Losers:       []uuid.UUID{},
		Drawn:        []uuid.UUID{},
		State:        StateNotStarted,
		mux:          &sync.Mutex{},
	}
	game.updatePub = NewGameUpdatesPub(game.GetID())

	for _, player := range r.PlayerOrder {
		timerRequest := timer.RequestNewTimer{
			StartingTimeMilis: r.PlayerTimeMilis,
			DecrementMilis:    1_000,
		}
		game.playerTimers[player] = timer.NewTimer(timerRequest)
	}

	gameStore := getGameStore()
	gameStore.Lock()
	defer gameStore.Unlock()

	gameStore.Store(game)

	return game, nil
}

// RequestGetGame is used to get a Game by its ID.
type RequestGetGame struct {
	GameID uuid.UUID `json:"game_id"`
}

// PerformAction loads a Game.
func (r *RequestGetGame) PerformAction() (*Game, errortypes.TypedError) {
	gameStore := getGameStore()
	gameStore.Lock()
	defer gameStore.Unlock()

	game := gameStore.GetByID(r.GameID)
	if game == nil {
		return nil, errGameNotFound{}
	}

	return game, nil
}

// RequestStartGame is used to start a Game.
type RequestStartGame struct {
	GameID uuid.UUID `json:"game_id"`
}

// PerformAction starts a Game.
func (r *RequestStartGame) PerformAction() (*Game, errortypes.TypedError) {
	e, err := (&RequestGetGame{GameID: r.GameID}).PerformAction()
	if err != nil {
		return nil, err
	}

	err = e.start()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// RequestConcedeGame is used to concede a Game.
type RequestConcedeGame struct {
	GameID   uuid.UUID `json:"game_id" mapstructure:"game_id" swaggerignore:"true"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestConcedeGame) PerformAction() (*Game, errortypes.TypedError) {
	game, err := (&RequestGetGame{GameID: r.GameID}).PerformAction()
	if err != nil {
		return nil, err
	}

	err = game.declareLoser(r.PlayerID)
	if err != nil {
		return nil, err
	}

	return game, nil
}
