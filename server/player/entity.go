package player

import "github.com/google/uuid"

// Player represents a user who interacts with other entities.
type Player struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"display_name"`
}

// GetID returns a Player's ID.
func (p Player) GetID() uuid.UUID {
	return p.ID
}
