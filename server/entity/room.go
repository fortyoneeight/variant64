package entity

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Room struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Players []uuid.UUID `json:"players"`

	mux *sync.RWMutex
}

func (r Room) GetID() uuid.UUID {
	return r.ID
}

type RequestNewRoom struct {
	Name string `json:"room_name"`
}

func (r *RequestNewRoom) Write(e *Entity[Room]) error {
	e.store = GetRoomStore()
	e.Data = &Room{
		ID:      uuid.New(),
		Name:    r.Name,
		Players: make([]uuid.UUID, 0),
		mux:     &sync.RWMutex{},
	}
	return nil
}

type RequestGetRoom struct {
	ID uuid.UUID `json:"room_id"`
}

func (r *RequestGetRoom) Read(e *Entity[Room]) error {
	e.store = GetRoomStore()
	e.Data = &Room{
		ID: r.ID,
	}
	return nil
}

type RequestGetRooms struct{}

func (r *RequestGetRooms) Read(e *EntityList[Room]) error {
	e.store = GetRoomStore()
	list := make([]*Room, 0)
	e.Data = &list
	return nil
}

type RequestRoomAddPlayer struct {
	RoomID   uuid.UUID `json:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestRoomAddPlayer) Write(e *Entity[Room]) error {
	getRoomReq := &RequestGetRoom{ID: r.RoomID}
	err := getRoomReq.Read(e)
	if err != nil {
		return err
	}
	err = e.Load()
	fmt.Println(e.Data)
	if err != nil {
		return err
	}

	player := &Entity[Player]{}
	getPlayerReq := &RequestGetPlayer{ID: r.PlayerID}
	err = getPlayerReq.Read(player)
	if err != nil {
		return err
	}
	err = player.Load()
	if err != nil {
		return err
	}

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

type RequestRoomRemovePlayer struct {
	RoomID   uuid.UUID `json:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestRoomRemovePlayer) Write(e *Entity[Room]) error {
	getRoomReq := &RequestGetRoom{ID: r.RoomID}
	err := getRoomReq.Read(e)
	if err != nil {
		return err
	}
	err = e.Load()
	if err != nil {
		return err
	}

	player := &Entity[Player]{}
	getPlayerReq := &RequestGetPlayer{ID: r.PlayerID}
	err = getPlayerReq.Read(player)
	if err != nil {
		return err
	}
	err = player.Load()
	if err != nil {
		return err
	}

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
