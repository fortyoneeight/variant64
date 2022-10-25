package entity

import (
	"github.com/google/uuid"
)

type Player struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"display_name"`
}

func (p Player) GetID() uuid.UUID {
	return p.ID
}

type RequestNewPlayer struct {
	DisplayName string `json:"display_name"`
}

func (r *RequestNewPlayer) Write(e *Entity[Player]) error {
	e.store = GetPlayerStore()
	e.Data = Player{
		ID: uuid.New(),
		DisplayName: r.DisplayName,
	}

	return nil
}

type RequestGetPlayer struct {
	ID uuid.UUID `json:"id"`
}

func (r *RequestGetPlayer) Read(e *Entity[Player]) error {
	e.store = GetPlayerStore()
	e.Data = Player{
		ID: r.ID,
	}

	return nil
}
