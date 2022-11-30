package room

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/errortypes"
	"github.com/variant64/server/game"
)

// RequestNewRoom is used to create a new Room.
type RequestNewRoom struct {
	Name string `json:"room_name" mapstructure:"room_name"`
}

// PerformAction creates a new Room.
func (r *RequestNewRoom) PerformAction() (*Room, errortypes.TypedError) {
	if r.Name == "" {
		return nil, errMissingName{}
	}

	room := &Room{
		ID:      uuid.New(),
		Name:    r.Name,
		Players: make([]uuid.UUID, 0),
		mux:     &sync.RWMutex{},
	}

	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	roomStore.Store(room)

	return room, nil
}

// RequestGetRoom is used to get a Room by its ID.
type RequestGetRoom struct {
	RoomID uuid.UUID `json:"room_id" mapstructure:"room_id"`
}

// PerformAction loads a Room.
func (r *RequestGetRoom) PerformAction() (*Room, errortypes.TypedError) {
	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	room := roomStore.GetByID(r.RoomID)
	if room == nil {
		return nil, errRoomNotFound{}
	}

	return room, nil
}

// RequestGetRooms is used to get all Rooms.
type RequestGetRooms struct{}

// PerformAction gets all rooms.
func (r *RequestGetRooms) PerformAction() ([]*Room, errortypes.TypedError) {
	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	rooms := roomStore.GetAll()
	if rooms == nil {
		return nil, errRoomNotFound{}
	}

	return rooms, nil
}

// RequestJoinRoom is used to add a Player to a Room.
type RequestJoinRoom struct {
	RoomID   uuid.UUID `json:"room_id" mapstructure:"room_id" swaggerignore:"true"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestJoinRoom) PerformAction() (*Room, errortypes.TypedError) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	for _, p := range room.Players {
		if p == r.PlayerID {
			return nil, errDuplicatePlayer{playerID: r.PlayerID}
		}
	}
	room.Players = append(room.Players, r.PlayerID)

	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	roomStore.Store(room)

	return room, nil
}

// RequestLeaveRoom is used to remove a Player from a Room.
type RequestLeaveRoom struct {
	RoomID   uuid.UUID `json:"room_id" mapstructure:"room_id" swaggerignore:"true"`
	PlayerID uuid.UUID `json:"player_id"`
}

// RequestLeaveRoom handles a RequestRoomAddPlayer.
func (r *RequestLeaveRoom) PerformAction() (*Room, errortypes.TypedError) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	for i, p := range room.Players {
		if p == r.PlayerID {
			room.Players = append(room.Players[:i], room.Players[i+1:]...)
		}
	}

	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	roomStore.Store(room)

	return room, nil
}

// RequestStartGame is used to start a new Game in a Room.
type RequestStartGame struct {
	RoomID          uuid.UUID `json:"room_id" mapstructure:"room_id"`
	PlayerTimeMilis int64     `json:"player_time_ms"`
}

// PerformAction starts a game.Game in a Room.
func (r *RequestStartGame) PerformAction() (*game.Game, errortypes.TypedError) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	gameEntity, err := (&game.RequestNewGame{
		PlayerOrder:     room.Players,
		PlayerTimeMilis: r.PlayerTimeMilis,
	}).PerformAction()
	if err != nil || gameEntity == nil {
		return nil, err
	}

	gameEntity, err = (&game.RequestStartGame{
		GameID: gameEntity.GetID(),
	}).PerformAction()
	if err != nil || gameEntity == nil {
		return nil, err
	}

	gameID := gameEntity.GetID()
	room.GameID = &gameID

	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	roomStore.Store(room)

	return gameEntity, nil
}
