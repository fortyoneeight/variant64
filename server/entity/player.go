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
	player := Player{
		ID: uuid.New(),
		DisplayName: r.DisplayName,
	}
	e.store = GetPlayerStore()
	e.Data = &player
	return nil
}

type RequestGetPlayer struct {
	ID uuid.UUID `json:"id"`
}

func (r *RequestGetPlayer) Read(e *Entity[Player]) error {
	player := Player{
		ID: r.ID,
	}
	e.store = GetPlayerStore()
	e.Data = &player
	return nil
}
