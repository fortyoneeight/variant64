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
	GameID  *uuid.UUID  `json:"game_id"`

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
	if r.Name == "" {
		return errors.New("room_name cannot be empty")
	}

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

	err := e.Load()
	if err != nil {
		return err
	}

	return nil
}

// RequestGetRooms is used to get all Rooms.
type RequestGetRooms struct{}

// Read adds all Rooms to the provided RoomList.
func (r *RequestGetRooms) Read(e *EntityList[Room]) {
	e.EntityStore = GetRoomStore()
	e.Data = make([]Room, 0)
	e.Load()
}
