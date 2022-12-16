package room

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/models"
)

const (
	PLAYER_LIMIT_DEFAULT = 2
	NAME_MAX_LENGTH = 16
)

// Room represents a room in which Players can play and watch Games.
type Room struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`

	Players     []uuid.UUID `json:"players"`
	PlayerLimit int         `json:"player_limit"`

	GameID *uuid.UUID `json:"game_id"`

	updateHandler *models.UpdatePublisher[RoomUpdate]

	mux *sync.RWMutex
}

// GetID returns a Room's ID.
func (r *Room) GetID() uuid.UUID {
	return r.ID
}

// RoomUpdate represents a change in a Room's state.
type RoomUpdate struct {
	ID      *uuid.UUID   `json:"id,omitempty"`
	Players *[]uuid.UUID `json:"players,omitempty"`
	GameID  *uuid.UUID   `json:"game_id,omitempty"`
}

// getSnapshot returns a RoomUpdate that is a snapshot of the Room.
func (r *Room) getSnapshot() RoomUpdate {
	return RoomUpdate{
		ID:      &r.ID,
		Players: &r.Players,
		GameID:  r.GameID,
	}
}
