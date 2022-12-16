package player

import (
	"github.com/google/uuid"
)

// RequestNewPlayer is used to create a new Player.
type RequestNewPlayer struct {
	DisplayName string `json:"display_name" mapstructure:"display_name"`
}

// PerformAction creates a new Player.
func (r *RequestNewPlayer) PerformAction() (*Player, error) {
	if r.DisplayName == "" {
		return nil, errMissingDisplayName
	}

	if len(r.DisplayName) > DISPLAY_NAME_MAX_LENGTH {
		return nil, errDisplayNameTooLong
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
		return nil, errPlayerNotFound
	}

	return player, nil
}
