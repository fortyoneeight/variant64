package entity

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/variant64/server/store"
)

type RequestHandler struct{}

type entityWriter[T store.Indexable] interface {
	Write(*Entity[T]) error
}

type entityReader[T store.Indexable] interface {
	Read(*Entity[T]) error
}

// HandleGetRooms handles a RequestGetRooms.
func (h *RequestHandler) HandleGetRooms(r *RequestGetRooms) *EntityList[Room] {
	rooms := &EntityList[Room]{}
	r.Read(rooms)
	return rooms
}

// RequestJoinRoom is used to add a Player to a Room.
type RequestJoinRoom struct {
	RoomID   uuid.UUID `json:"room_id" mapstructure:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestJoinRoom) PerformAction() (*Entity[Room], error) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	for _, p := range room.Data.Players {
		if p == r.PlayerID {
			return nil, errors.New("player cannot be duplicate")
		}
	}
	room.Data.Players = append(room.Data.Players, r.PlayerID)

	room.Store()

	return room, nil
}

// RequestLeaveRoom is used to remove a Player from a Room.
type RequestLeaveRoom struct {
	RoomID   uuid.UUID `json:"room_id" mapstructure:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

// RequestLeaveRoom handles a RequestRoomAddPlayer.
func (r *RequestLeaveRoom) PerformAction() (*Entity[Room], error) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	for i, p := range room.Data.Players {
		if p == r.PlayerID {
			room.Data.Players = append(room.Data.Players[:i], room.Data.Players[i+1:]...)
		}
	}

	room.Store()

	return room, nil
}

// RequestStartGame is used to start a new Game in a Room.
type RequestStartGame struct {
	RoomID          uuid.UUID `json:"room_id" mapstructure:"room_id"`
	PlayerTimeMilis int64     `json:"player_time_ms"`
}

// HandleRequestGameStart handles a RequestGameStart.
func (r *RequestStartGame) PerformAction() (*Entity[Room], error) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	requestNewGame := &RequestNewGame{
		PlayerOrder:     room.Data.Players,
		PlayerTimeMilis: r.PlayerTimeMilis,
	}

	game, err := requestNewGame.PerformAction()
	if err != nil || game == nil {
		return nil, err
	}

	gameID := game.Data.GetID()
	room.Data.GameID = &gameID

	game.Data.start()

	room.Store()
	game.Store()

	return room, nil
}
