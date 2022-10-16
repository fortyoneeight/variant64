package room

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/player"
)

type Room struct {
	ID      uuid.UUID        `json:"id"`
	Name    string           `json:"name"`
	Players []player.Player `json:"players"`

	mux *sync.RWMutex
}

func NewRoom(name string) *Room {
	return &Room{
		ID:   uuid.New(),
		Name: name,
		Players: make([]player.Player, 0),
		mux:  &sync.RWMutex{},
	}
}

func (r *Room) GetID() uuid.UUID {
	return r.ID
}

func (r *Room) AddPlayer(newPlayer player.Player) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, p := range r.Players {
		if p.ID == newPlayer.ID {
			return errors.New("player cannot be duplicate")
		}
	}
	r.Players = append(r.Players, newPlayer)
	return nil
}

func (r *Room) RemovePlayer(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	for i, p := range r.Players {
		if p.ID == id {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			return nil
		}
	}
	return errors.New("player not found")
}
