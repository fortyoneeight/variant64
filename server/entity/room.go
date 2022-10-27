package entity

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// Room represents a room in which Players can play and watch Games.
type Room struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Players []uuid.UUID `json:"players"`

	mux *sync.RWMutex
}

// GetID returns a Room's ID.
func (r Room) GetID() uuid.UUID {
	return r.ID
}

// RequestNewRoom is used to create a new Room.
type RequestNewRoom struct {
	Name string `json:"room_name"`
}

// Write initializes all fields of the provided Room.
func (r *RequestNewRoom) Write(e *Entity[Room]) error {
	e.EntityStore = GetRoomStore()
	e.Data = Room{
		ID:      uuid.New(),
		Name:    r.Name,
		Players: make([]uuid.UUID, 0),
		mux:     &sync.RWMutex{},
	}
	return nil
}

// RequestGetRoom is used to get a Room by its ID.
type RequestGetRoom struct {
	ID uuid.UUID `json:"room_id"`
}

// Read intializes the ID field of the provided Room.
func (r *RequestGetRoom) Read(e *Entity[Room]) error {
	e.EntityStore = GetRoomStore()
	e.Data = Room{
		ID: r.ID,
	}
	return nil
}

// RequestGetRooms is used to get all Rooms.
type RequestGetRooms struct{}

// Read adds all Rooms to the provided RoomList.
func (r *RequestGetRooms) Read(e *EntityList[Room]) error {
	e.EntityStore = GetRoomStore()
	e.Data = make([]Room, 0)
	return nil
}

// RequestRoomAddPlayer is used to add a Player to a Room.
type RequestRoomAddPlayer struct {
	RoomID   uuid.UUID `json:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

// Write adds the Request's Player to the provided Room.
func (r *RequestRoomAddPlayer) Write(e *Entity[Room]) error {
	e.Data.mux.Lock()
	defer e.Data.mux.Unlock()

	for _, p := range e.Data.Players {
		if p == r.PlayerID {
			return errors.New("player cannot be duplicate")
		}
	}
	e.Data.Players = append(e.Data.Players, r.PlayerID)
	return nil
}

// RequestRoomRemovePlayer is used to remove a Player from a Room.
type RequestRoomRemovePlayer struct {
	RoomID   uuid.UUID `json:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

// Write removes the Request's Player from the provided Room.
func (r *RequestRoomRemovePlayer) Write(e *Entity[Room]) error {
	e.Data.mux.Lock()
	defer e.Data.mux.Unlock()

	for i, p := range e.Data.Players {
		if p == r.PlayerID {
			e.Data.Players = append(e.Data.Players[:i], e.Data.Players[i+1:]...)
			return nil
		}
	}
	return nil
}
