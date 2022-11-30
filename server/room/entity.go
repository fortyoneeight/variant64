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

	mux *sync.RWMutex
}

// GetID returns a Room's ID.
func (r Room) GetID() uuid.UUID {
	return r.ID
}
