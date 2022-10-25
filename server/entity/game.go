package entity

import (
	"errors"
	"github.com/google/uuid"
)

// Game represents an on-going game between a list of players.
type Game struct {
	id uuid.UUID

	activePlayer uuid.UUID
	playerTimers map[uuid.UUID]*Timer
	playerOrder  []uuid.UUID

	updatePub *updatePub[GameUpdate]
}

// GameUpdate represents a change in a Game's state.
type GameUpdate struct {
	GameID uuid.UUID `json:"game_id"`

	ActivePlayer *uuid.UUID           `json:"active_player,omitempty"`
	Clocks       *map[uuid.UUID]int64 `json:"clocks,omitempty"`
}

// GetID returns a Game's ID.
func (g Game) GetID() uuid.UUID {
	return g.id
}

// start initializes the game and starts it.
func (g *Game) start() {
	g.subscribeToTimers()
	for _, timer := range g.playerTimers {
		timer.Start()
	}
	g.playerTimers[g.activePlayer].Unpause()
}

// passTurn passes the turn to the next player
// the active player's clock pauses and the next player's clock unpauses.
func (g *Game) passTurn() {
	g.playerTimers[g.activePlayer].Pause()
	g.playerTimers[g.playerOrder[0]].Unpause()

	g.activePlayer = g.playerOrder[0]
	g.playerOrder = append(g.playerOrder[1:], g.playerOrder[0])
}

// subscribeToTimers starts a subscription to each player's Timer.
func (g *Game) subscribeToTimers() {
	for playerID, timer := range g.playerTimers {
		go g.handleTimerUpdate(playerID, timer)
	}
}

// handleTimerUpdate publishes a GameUpdate every time a player's Timer updates.
func (g *Game) handleTimerUpdate(playerID uuid.UUID, t *Timer) {
	for {
		select {
		case val := <-t.TimerChan:
			g.updatePub.Publish(
				GameUpdate{
					GameID: g.id,
					Clocks: &map[uuid.UUID]int64{playerID: val},
				},
			)
		}
	}
}

// RequestNewGame is a used to create a new Game.
type RequestNewGame struct {
	PlayerOrder     []uuid.UUID `json:"player_order"`
	PlayerTimeMilis int64       `json:"player_time_ms"`
}

// Write initializes all fields of the provided Game.
func (r *RequestNewGame) Write(e *Entity[*Game]) error {
	if len(r.PlayerOrder) < 2 {
		return errors.New("invalid number of players, must be >= 2")
	}

	e.EntityStore = GetGameStore()
	e.Data = &Game{
		id:           uuid.New(),
		playerOrder:  append(r.PlayerOrder[1:], r.PlayerOrder[0]),
		playerTimers: make(map[uuid.UUID]*Timer),
		activePlayer: r.PlayerOrder[0],
	}
	e.Data.updatePub = NewGameUpdatesPub(e.Data.id)

	for _, player := range r.PlayerOrder {
		timerRequest := RequestNewTimer{
			StartingTimeMilis: r.PlayerTimeMilis,
			DecrementMilis:    1_000,
		}
		e.Data.playerTimers[player] = NewTimer(timerRequest)
	}

	return nil
}

// RequestGetGame is used to get a Game by its ID.
type RequestGetGame struct {
	ID uuid.UUID `json:"game_id"`
}

// Read intializes the ID field of the provided Game.
func (r *RequestGetGame) Read(e *Entity[*Game]) error {
	e.EntityStore = GetGameStore()
	e.Data = &Game{
		id: r.ID,
	}
	return nil
}

// RequestGameStart is used to start a Game.
type RequestGameStart struct{}

// Write starts the provided Game.
func (r *RequestGameStart) Write(e *Entity[*Game]) error {
	e.Data.start()
	return nil
}

// RequestGamePassTurn is used to pass the turn in a Game.
type RequestGamePassTurn struct{}

// Write passes the turn to the next player of the provided Game.
func (r *RequestGamePassTurn) Write(e *Entity[*Game]) error {
	e.Data.passTurn()
	return nil
}
