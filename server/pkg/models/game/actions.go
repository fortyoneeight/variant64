package game

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/errortypes"
	"github.com/variant64/server/pkg/models"
	"github.com/variant64/server/pkg/timer"

	"github.com/gorilla/websocket"
)

var gameUpdateBus = models.NewUpdateBus[GameUpdate]()

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
		ApprovedDraw: map[uuid.UUID]bool{},
		State:        StateNotStarted,
		mux:          &sync.Mutex{},
	}
	handler, err := models.NewUpdatePub(game.ID, gameUpdateBus)
	if err != nil {
		return nil, models.ErrFailedUpdatePub{}
	}

	game.updateHandler = handler

	for _, player := range r.PlayerOrder {
		game.ApprovedDraw[player] = false

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

// RequestConcede is used to concede a Game.
type RequestConcede struct {
	GameID   uuid.UUID `json:"game_id" mapstructure:"game_id" swaggerignore:"true"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestConcede) PerformAction() (*Game, errortypes.TypedError) {
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

// RequestApproveDraw is used to approve a draw for a Game.
type RequestApproveDraw struct {
	GameID   uuid.UUID `json:"game_id" mapstructure:"game_id" swaggerignore:"true"`
	PlayerID uuid.UUID `json:"player_id"`
}

// PerformAction approves a draw for one player in a Game.
func (r *RequestApproveDraw) PerformAction() (*Game, errortypes.TypedError) {
	game, err := (&RequestGetGame{GameID: r.GameID}).PerformAction()
	if err != nil {
		return nil, err
	}

	err = game.approveDraw(r.PlayerID)
	if err != nil {
		return nil, err
	}

	return game, nil
}

// RequestRejectDraw is used to reject a draw for a Game.
type RequestRejectDraw struct {
	GameID uuid.UUID `json:"game_id" mapstructure:"game_id" swaggerignore:"true"`
}

// PerformAction rejects a draw in a Game.
func (r *RequestRejectDraw) PerformAction() (*Game, errortypes.TypedError) {
	game, err := (&RequestGetGame{GameID: r.GameID}).PerformAction()
	if err != nil {
		return nil, err
	}

	err = game.rejectDraw()
	if err != nil {
		return nil, err
	}

	return game, nil
}

const (
	MessageChannel = "game"

	GameSubscribe   string = "subscribe"
	GameUnsubscribe string = "unsubscribe"
)

// CommandGameSubscribe represents a game subscribe command.
type CommandGameSubscribe struct {
	GameID      uuid.UUID `json:"game_id"`
	EventWriter models.EventWriter
}

func (c *CommandGameSubscribe) PerformAction() errortypes.TypedError {
	models.Subscribe(gameUpdateBus, c.GameID, MessageChannel, c.EventWriter)
	return nil
}

// CommandGameUnsubscribe represents an game unsubscribe command.
type CommandGameUnsubscribe struct {
	models.Command
	GameID uuid.UUID `json:"game_id"`
}

func (c *CommandGameUnsubscribe) PerformAction(conn *websocket.Conn) errortypes.TypedError {
	return nil
}

// HandleCommand handles all incoming game writer messages.
func HandleCommand(writer models.EventWriter, command, body string) errortypes.TypedError {
	switch {
	case command == GameSubscribe:
		return models.HandleCommand(models.MarshallCommand(body, &CommandGameSubscribe{EventWriter: writer}))
	default:
		return models.ErrInvalidCommand{}
	}
}
