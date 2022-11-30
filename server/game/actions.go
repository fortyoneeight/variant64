package game

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/variant64/server/entity"
	"github.com/variant64/server/timer"
)

// RequestNewGame is a used to create a new Game.
type RequestNewGame struct {
	PlayerOrder     []uuid.UUID `json:"player_order" swaggerignore:"true"`
	PlayerTimeMilis int64       `json:"player_time_ms"`
}

// PerformAction creates a new Game.
func (r *RequestNewGame) PerformAction() (*entity.Entity[*Game], error) {
	if len(r.PlayerOrder) < 2 {
		return nil, errors.New("invalid number of players, must be >= 2")
	}

	e := &entity.Entity[*Game]{}
	e.EntityStore = getGameStore()
	e.Data = &Game{
		id:           uuid.New(),
		playerOrder:  append(r.PlayerOrder[1:], r.PlayerOrder[0]),
		playerTimers: make(map[uuid.UUID]*timer.Timer),
		activePlayer: r.PlayerOrder[0],
	}
	e.Data.updatePub = NewGameUpdatesPub(e.Data.id)

	for _, player := range r.PlayerOrder {
		timerRequest := timer.RequestNewTimer{
			StartingTimeMilis: r.PlayerTimeMilis,
			DecrementMilis:    1_000,
		}
		e.Data.playerTimers[player] = timer.NewTimer(timerRequest)
	}

	e.Store()

	return e, nil
}

// RequestGetGame is used to get a Game by its ID.
type RequestGetGame struct {
	GameID uuid.UUID `json:"game_id"`
}

// PerformAction loads a Game.
func (r *RequestGetGame) PerformAction() (*entity.Entity[*Game], error) {
	e := &entity.Entity[*Game]{}
	e.EntityStore = getGameStore()
	e.Data = &Game{
		id: r.GameID,
	}

	err := e.Load()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// RequestGamePassTurn is used to pass the turn in a Game.
type RequestStartGame struct {
	GameID uuid.UUID `json:"game_id"`
}

// PerformAction starts a Game.
func (r *RequestStartGame) PerformAction() (*entity.Entity[*Game], error) {
	e, err := (&RequestGetGame{GameID: r.GameID}).PerformAction()
	if err != nil {
		return nil, err
	}

	e.Data.start()

	return e, nil
}

// RequestGamePassTurn is used to pass the turn in a Game.
type RequestGamePassTurn struct{}

// Write passes the turn to the next player of the provided Game.
func (r *RequestGamePassTurn) Write(e *entity.Entity[*Game]) error {
	e.Data.passTurn()
	return nil
}
