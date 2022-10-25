package entity

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Game struct {
	id uuid.UUID

	activePlayer uuid.UUID
	playerTimers map[uuid.UUID]*Timer
	playerOrder  []uuid.UUID

	subscribers []GameSubscription
}

type GameUpdate struct {
	ActivePlayer uuid.UUID           `json:"active_player"`
	Clocks       map[uuid.UUID]int64 `json:"clocks"`
}

// TimerSubscription is used to subscribe to Timer updates
type GameSubscription interface {
	OnUpdate(u GameUpdate)
}

func (g Game) GetID() uuid.UUID {
	return g.id
}

// start initializes the game and starts it
func (g *Game) start() {
	g.subscribeToTimers()
	for _, timer := range g.playerTimers {
		timer.Start()
	}
	g.playerTimers[g.activePlayer].Unpause()
}

// passTurn passes the turn to the next player
// the active player's clock pauses and the next player's clock unpauses
func (g *Game) passTurn() {
	g.playerTimers[g.activePlayer].Pause()
	g.playerTimers[g.playerOrder[0]].Unpause()

	g.activePlayer = g.playerOrder[0]
	g.playerOrder = append(g.playerOrder[1:], g.playerOrder[0])
}

// subscribe adds the subscription to the list of subscribers
func (g *Game) subscribe(sub GameSubscription) {
	g.subscribers = append(g.subscribers, sub)
}

// publishUpdate sends a GameUpdate to all subscribers
func (g *Game) publishUpdate(update GameUpdate) {
	for _, sub := range g.subscribers {
		sub.OnUpdate(update)
	}
}

// subscribeToTimers starts a subscription to each player's Timer
func (g *Game) subscribeToTimers() {
	for playerID, timer := range g.playerTimers {
		go g.handleTimerUpdate(playerID, timer)
	}
}

// handleTimerUpdate publishes a GameUpdate every time a player's Timer updates
func (g *Game) handleTimerUpdate(playerID uuid.UUID, t *Timer) {
	for {
		select {
		case val := <-t.TimerChan:
			fmt.Println("update")
			g.publishUpdate(
				GameUpdate{
					Clocks: map[uuid.UUID]int64{playerID: val},
				},
			)
		}
	}
}

type RequestNewGame struct {
	PlayerOrder     []uuid.UUID `json:"player_order"`
	PlayerTimeMilis int64       `json:"player_time_ms"`
}

func (r *RequestNewGame) Write(e *Entity[*Game]) error {
	if len(r.PlayerOrder) < 2 {
		return errors.New("invalid number of players, must be >= 2")
	}

	e.Data.playerOrder = r.PlayerOrder
	e.Data.playerOrder = append(e.Data.playerOrder[1:], e.Data.playerOrder[0])
	e.Data.playerTimers = make(map[uuid.UUID]*Timer)
	e.Data.activePlayer = r.PlayerOrder[0]
	e.Data.subscribers = make([]GameSubscription, 0)

	for _, player := range r.PlayerOrder {
		timerRequest := RequestNewTimer{
			StartingTimeMilis: r.PlayerTimeMilis,
			DecrementMilis:    1_000,
		}
		e.Data.playerTimers[player] = NewTimer(timerRequest)
	}

	return nil
}

type RequestGetGame struct {
	ID uuid.UUID `json:"game_id"`
}

func (r *RequestGetGame) Read(e *Entity[*Game]) error {
	e.store = GetGameStore()
	e.Data = &Game{
		id: r.ID,
	}
	return nil
}

type RequestGameAddSubsciption struct {
	subscriber GameSubscription
}

func (r *RequestGameAddSubsciption) Write(e *Entity[Game]) error {
	e.Data.subscribers = append(e.Data.subscribers, r.subscriber)
	return nil
}

type RequestGamePassTurn struct{}

func (r *RequestGamePassTurn) Write(e *Entity[Game]) error {
	e.Data.passTurn()
	return nil
}
