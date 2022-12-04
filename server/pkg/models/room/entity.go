package room

import (
	"sync"

	"github.com/google/uuid"
)

// Room represents a room in which Players can play and watch Games.
type Room struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Players []uuid.UUID `json:"players"`
	GameID  *uuid.UUID  `json:"game_id"`

	updatePub *updatePub[RoomUpdate]

	mux *sync.RWMutex
}

// GetID returns a Room's ID.
func (r *Room) GetID() uuid.UUID {
	return r.ID
}

// RoomUpdate represents a change in a Room's state.
type RoomUpdate struct {
	ID      uuid.UUID   `json:"id,omitempty"`
	Players []uuid.UUID `json:"players,omitempty"`
	GameID  *uuid.UUID  `json:"game_id,omitempty"`
}
