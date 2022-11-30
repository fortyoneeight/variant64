package player

import (
	"errors"

	"github.com/google/uuid"
)

// RequestNewPlayer is used to create a new Player.
type RequestNewPlayer struct {
	DisplayName string `json:"display_name" mapstructure:"display_name"`
}

// PerformAction creates a new Player.
func (r *RequestNewPlayer) PerformAction() (*Player, error) {
	if r.DisplayName == "" {
		return nil, errors.New("display_name cannot be empty")
	}

	player := &Player{
		ID:          uuid.New(),
		DisplayName: r.DisplayName,
	}

	playerStore := getPlayerStore()
	playerStore.Lock()
	defer playerStore.Unlock()

	playerStore.Store(player)

	return player, nil
}

// RequestGetPlayer is used to get a Player by its ID.
type RequestGetPlayer struct {
	PlayerID uuid.UUID `json:"player_id" mapstructure:"player_id"`
}

// PerformAction loads a Player.
func (r *RequestGetPlayer) PerformAction() (*Player, error) {
	playerStore := getPlayerStore()
	playerStore.Lock()
	defer playerStore.Unlock()

	player := playerStore.GetByID(r.PlayerID)
	if player == nil {
		return nil, errors.New("not found")
	}

	return player, nil
}
