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
	DisplayName string `json:"display_name" mapstructure:"display_name"`
}

// PerformAction creates a new Player.
func (r *RequestNewPlayer) PerformAction() (*Entity[Player], error) {
	e := &Entity[Player]{}
	if r.DisplayName == "" {
		return nil, errors.New("display_name cannot be empty")
	}

	e.EntityStore = GetPlayerStore()
	e.Data = Player{
		ID:          uuid.New(),
		DisplayName: r.DisplayName,
	}

	e.Store()

	return e, nil
}

// RequestGetPlayer is used to get a Player by its ID.
type RequestGetPlayer struct {
	PlayerID uuid.UUID `json:"player_id" mapstructure:"player_id"`
}

// PerformAction loads a Player.
func (r *RequestGetPlayer) PerformAction() (*Entity[Player], error) {
	e := &Entity[Player]{}
	e.EntityStore = GetPlayerStore()
	e.Data = Player{
		ID: r.PlayerID,
	}

	err := e.Load()
	if err != nil {
		return nil, err
	}

	return e, nil
}
