package game

import (
	"github.com/google/uuid"
	"github.com/variant64/server/pkg/timer"
)

// Game represents an on-going game between a list of players.
type Game struct {
	id uuid.UUID

	activePlayer uuid.UUID
	playerTimers map[uuid.UUID]*timer.Timer
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
func (g *Game) handleTimerUpdate(playerID uuid.UUID, t *timer.Timer) {
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
