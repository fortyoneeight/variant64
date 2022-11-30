package room

import (
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/variant64/server/entity"
	"github.com/variant64/server/game"
)

// RequestNewRoom is used to create a new Room.
type RequestNewRoom struct {
	Name string `json:"room_name" mapstructure:"room_name"`
}

// PerformAction creates a new Room.
func (r *RequestNewRoom) PerformAction() (*entity.Entity[Room], error) {
	if r.Name == "" {
		return nil, errors.New("room_name cannot be empty")
	}

	e := &entity.Entity[Room]{}
	e.EntityStore = getRoomStore()
	e.Data = Room{
		ID:      uuid.New(),
		Name:    r.Name,
		Players: make([]uuid.UUID, 0),
		mux:     &sync.RWMutex{},
	}

	e.Store()

	return e, nil
}

// RequestGetRoom is used to get a Room by its ID.
type RequestGetRoom struct {
	RoomID uuid.UUID `json:"room_id" mapstructure:"room_id"`
}

// PerformAction loads a Room.
func (r *RequestGetRoom) PerformAction() (*entity.Entity[Room], error) {
	e := &entity.Entity[Room]{}
	e.EntityStore = getRoomStore()
	e.Data = Room{
		ID: r.RoomID,
	}

	err := e.Load()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// RequestGetRooms is used to get all Rooms.
type RequestGetRooms struct{}

// Read adds all Rooms to the provided RoomList.
func (r *RequestGetRooms) Read(e *entity.EntityList[Room]) {
	e.EntityStore = getRoomStore()
	e.Data = make([]Room, 0)
	e.Load()
}

// RequestJoinRoom is used to add a Player to a Room.
type RequestJoinRoom struct {
	RoomID   uuid.UUID `json:"room_id" mapstructure:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestJoinRoom) PerformAction() (*entity.Entity[Room], error) {
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
func (r *RequestLeaveRoom) PerformAction() (*entity.Entity[Room], error) {
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
func (r *RequestStartGame) PerformAction() (*entity.Entity[Room], error) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	gameEntity, err := (&game.RequestNewGame{
		PlayerOrder:     room.Data.Players,
		PlayerTimeMilis: r.PlayerTimeMilis,
	}).PerformAction()
	if err != nil || gameEntity == nil {
		return nil, err
	}

	gameEntity, err = (&game.RequestStartGame{
		GameID: gameEntity.Data.GetID(),
	}).PerformAction()
	if err != nil || gameEntity == nil {
		return nil, err
	}

	gameID := gameEntity.Data.GetID()
	room.Data.GameID = &gameID

	room.Store()

	return room, nil
}

type RequestHandler struct{}

// HandleGetRooms handles a RequestGetRooms.
func (h *RequestHandler) HandleGetRooms(r *RequestGetRooms) *entity.EntityList[Room] {
	rooms := &entity.EntityList[Room]{}
	r.Read(rooms)
	return rooms
}
