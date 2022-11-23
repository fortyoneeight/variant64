package entity

import (
	"errors"

	"github.com/google/uuid"
)

// Player represents a user who interacts with other entities.
type Player struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"display_name"`
}

// GetID returns a Player's ID.
func (p Player) GetID() uuid.UUID {
	return p.ID
}

// RequestNewPlayer is used to create a new Player.
type RequestNewPlayer struct {
	DisplayName string `json:"display_name"`
}

// Write initializes all fields of the provided Player.
func (r *RequestNewPlayer) Write(e *Entity[Player]) error {
	if r.DisplayName == "" {
		return errors.New("display_name cannot be empty")
	}

	e.EntityStore = GetPlayerStore()
	e.Data = Player{
		ID:          uuid.New(),
		DisplayName: r.DisplayName,
	}

	return nil
}

// RequestGetPlayer is used to get a Player by its ID.
type RequestGetPlayer struct {
	ID uuid.UUID `json:"id"`
}

// Read intializes the ID field of the provided Player.
func (r *RequestGetPlayer) Read(e *Entity[Player]) error {
	e.EntityStore = GetPlayerStore()
	e.Data = Player{
		ID: r.ID,
	}

	return nil
}
