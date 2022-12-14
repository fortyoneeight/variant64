package game

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/models"
	"github.com/variant64/server/pkg/models/board"
	"github.com/variant64/server/pkg/timer"
)

type gameState string

const (
	StateNotStarted gameState = "not_started"
	StateStarted    gameState = "started"
	StateFinished   gameState = "finished"
)

type gameboard interface {
	GetState() board.GameboardState
	HandleMove(move board.Move) error
}

// Game represents an on-going game between a list of players.
type Game struct {
	ID uuid.UUID `json:"id"`

	ActivePlayer uuid.UUID           `json:"active_player"`
	Clocks       map[uuid.UUID]int64 `json:"clocks,omitempty"`
	playerTimers map[uuid.UUID]*timer.Timer
	playerOrder  []uuid.UUID

	Winners      []uuid.UUID        `json:"winning_players"`
	Losers       []uuid.UUID        `json:"losing_players"`
	Drawn        []uuid.UUID        `json:"drawn_players"`
	ApprovedDraw map[uuid.UUID]bool `json:"approved_draw_players"`

	State gameState `json:"state"`

	board gameboard

	updateHandler *models.UpdatePublisher[GameUpdate]

	mux *sync.RWMutex
}

// GameUpdate represents a change in a Game's state.
type GameUpdate struct {
	ID uuid.UUID `json:"id"`

	ActivePlayer *uuid.UUID           `json:"active_player,omitempty"`
	Clocks       *map[uuid.UUID]int64 `json:"clocks,omitempty"`

	Winners *[]uuid.UUID `json:"winning_players,omitempty"`
	Losers  *[]uuid.UUID `json:"losing_players,omitempty"`
	Drawn   *[]uuid.UUID `json:"drawn_players,omitempty"`

	ApprovedDraw *map[uuid.UUID]bool `json:"approved_draw_players,omitempty"`

	State *gameState `json:"state,omitempty"`

	BoardState board.GameboardState `json:"gameboard_state,omitempty"`
}

// Build returns a GameUpdate.
func (g *GameUpdate) Build() GameUpdate {
	return *g
}

// GetID returns a Game's ID.
func (g *Game) GetID() uuid.UUID {
	return g.ID
}

// start initializes the game and starts it.
func (g *Game) start() error {
	g.mux.Lock()
	defer g.mux.Unlock()

	err := g.isGameInState(StateNotStarted)
	if err != nil {
		return err
	}

	g.subscribeToTimers()
	for _, timer := range g.playerTimers {
		timer.Start()
	}
	g.playerTimers[g.ActivePlayer].Unpause()
	g.State = StateStarted

	g.updateHandler.Publish(
		models.UpdateMessage[GameUpdate]{
			Channel: MessageChannel,
			Type:    models.UpdateType_DELTA,
			Data: GameUpdate{
				ID:           g.ID,
				ActivePlayer: &g.ActivePlayer,
				State:        &g.State,
				BoardState:   nil,
			},
		},
	)

	return nil
}

// passTurn passes the turn to the next player
// the active player's clock pauses and the next player's clock unpauses.
func (g *Game) passTurn() {
	g.playerTimers[g.ActivePlayer].Pause()
	g.playerTimers[g.playerOrder[0]].Unpause()

	g.ActivePlayer = g.playerOrder[0]
	g.playerOrder = append(g.playerOrder[1:], g.playerOrder[0])
}

// declareLoser sets one player as the loser and all other players as winners.
func (g *Game) declareLoser(playerID uuid.UUID) error {
	g.mux.Lock()
	defer g.mux.Unlock()

	err := g.isGameInState(StateStarted)
	if err != nil {
		return err
	}

	winners := make([]uuid.UUID, 0)
	losers := make([]uuid.UUID, 0)
	for _, p := range g.playerOrder {
		if p == playerID {
			losers = append(losers, p)
		} else {
			winners = append(winners, p)
		}
	}

	if len(losers) == 0 {
		return errPlayerNotInGame
	}

	g.playerTimers[g.ActivePlayer].Pause()
	g.Winners = winners
	g.Losers = losers
	g.State = StateFinished

	g.updateHandler.Publish(
		models.UpdateMessage[GameUpdate]{
			Channel: MessageChannel,
			Type:    models.UpdateType_DELTA,
			Data: GameUpdate{
				ID:      g.ID,
				Winners: &g.Winners,
				Losers:  &g.Losers,
				Drawn:   &g.Drawn,
				State:   &g.State,
			},
		},
	)

	return nil
}

// approveDraw marks the player as having accepeted a draw,
// if all players have accepeted a draw the game is considered a draw.
func (g *Game) approveDraw(playerID uuid.UUID) error {
	g.mux.Lock()
	defer g.mux.Unlock()

	err := g.isGameInState(StateStarted)
	if err != nil {
		return err
	}

	if _, ok := g.ApprovedDraw[playerID]; ok {
		g.ApprovedDraw[playerID] = true

		allAccepted := true
		for _, accepted := range g.ApprovedDraw {
			if !accepted {
				allAccepted = false
				break
			}
		}

		if allAccepted {
			g.Drawn = g.playerOrder
			g.State = StateFinished
			g.updateHandler.Publish(
				models.UpdateMessage[GameUpdate]{
					Channel: MessageChannel,
					Type:    models.UpdateType_DELTA,
					Data: GameUpdate{
						ID:      g.ID,
						Winners: &g.Winners,
						Losers:  &g.Losers,
						Drawn:   &g.Drawn,
						State:   &g.State,
					},
				},
			)
		} else {
			g.updateHandler.Publish(
				models.UpdateMessage[GameUpdate]{
					Channel: MessageChannel,
					Type:    models.UpdateType_DELTA,
					Data: GameUpdate{
						ID:      g.ID,
						Winners: &g.Winners,
						Losers:  &g.Losers,
						Drawn:   &g.Drawn,
						State:   &g.State,
					},
				},
			)
		}
	} else {
		return errPlayerNotInGame
	}

	return nil
}

// rejectDraw marks all the players as having not accepted a draw.
func (g *Game) rejectDraw() error {
	g.mux.Lock()
	defer g.mux.Unlock()

	err := g.isGameInState(StateStarted)
	if err != nil {
		return err
	}

	for player := range g.ApprovedDraw {
		g.ApprovedDraw[player] = false
	}

	g.updateHandler.Publish(
		models.UpdateMessage[GameUpdate]{
			Channel: MessageChannel,
			Type:    models.UpdateType_DELTA,
			Data: GameUpdate{
				ID:           g.ID,
				ApprovedDraw: &g.ApprovedDraw,
			},
		},
	)

	return nil
}

// makeMove makes a move for the provided PlayerID if valid.
func (g *Game) makeMove(playerID uuid.UUID, move board.Move) error {
	g.mux.Lock()
	defer g.mux.Unlock()

	err := g.isGameInState(StateStarted)
	if err != nil {
		return err
	}

	if g.ActivePlayer != playerID {
		return errNotPlayersTurn(playerID.String())
	}

	moveErr := g.board.HandleMove(move)
	if moveErr != nil {
		return errInvalidMove(err)
	}

	g.passTurn()

	g.updateHandler.Publish(
		models.UpdateMessage[GameUpdate]{
			Channel: MessageChannel,
			Type:    models.UpdateType_DELTA,
			Data: GameUpdate{
				ID: g.ID,
			},
		},
	)

	return nil
}

// isGameInState checks if the Game is in the correct state.
func (g *Game) isGameInState(required gameState) error {
	if g.State != required {
		return errIncorrectGameState(required, g.State)
	}
	return nil
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
			g.mux.Lock()

			g.Clocks[playerID] = val

			g.updateHandler.Publish(
				models.UpdateMessage[GameUpdate]{
					Channel: MessageChannel,
					Type:    models.UpdateType_DELTA,
					Data: GameUpdate{
						ID:     g.ID,
						Clocks: &map[uuid.UUID]int64{playerID: val},
					},
				},
			)

			g.mux.Unlock()
		}
	}
}

// getSnapshot returns a snapshot of the game state.
func (g *Game) getSnapshot() GameUpdate {
	g.mux.RLock()
	defer g.mux.RUnlock()

	return GameUpdate{
		ID:           g.ID,
		ActivePlayer: &g.ActivePlayer,
		Clocks:       &g.Clocks,
		Winners:      &g.Winners,
		Losers:       &g.Losers,
		Drawn:        &g.Drawn,
		ApprovedDraw: &g.ApprovedDraw,
		State:        &g.State,
		BoardState:   g.board.GetState(),
	}
}
