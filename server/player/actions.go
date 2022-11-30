package player

import (
	"errors"

	"github.com/google/uuid"
	"github.com/variant64/server/entity"
)

// RequestNewPlayer is used to create a new Player.
type RequestNewPlayer struct {
	DisplayName string `json:"display_name" mapstructure:"display_name"`
}

// PerformAction creates a new Player.
func (r *RequestNewPlayer) PerformAction() (*entity.Entity[Player], error) {
	e := &entity.Entity[Player]{}
	if r.DisplayName == "" {
		return nil, errors.New("display_name cannot be empty")
	}

	e.EntityStore = getPlayerStore()
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
func (r *RequestGetPlayer) PerformAction() (*entity.Entity[Player], error) {
	e := &entity.Entity[Player]{}
	e.EntityStore = getPlayerStore()
	e.Data = Player{
		ID: r.PlayerID,
	}

	err := e.Load()
	if err != nil {
		return nil, err
	}

	return e, nil
}
