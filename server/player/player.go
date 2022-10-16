package player

import "github.com/google/uuid"

type Player struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"display_name"`
}

func NewPlayer(displayName string) *Player {
	return &Player{
		ID:          uuid.New(),
		DisplayName: displayName,
	}
}

func (p *Player) GetID() uuid.UUID {
	return p.ID
}
