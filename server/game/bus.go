package game

import (
	"github.com/google/uuid"
	"github.com/variant64/server/bus"
)

type updatePub[T any] struct {
	pub      *bus.Pub[T]
	entityID uuid.UUID
}

func (u *updatePub[T]) Publish(update T) {
	u.pub.Publish(u.entityID, update)
}

var gameUpdateBus *bus.Bus[GameUpdate]

// GetGameUpdateBus returns the global bus for Game updates.
func GetGameUpdateBus() *bus.Bus[GameUpdate] {
	if gameUpdateBus == nil {
		gameUpdateBus = bus.NewBus[GameUpdate]([]uuid.UUID{})
		gameUpdateBus.Start()
	}
	return gameUpdateBus
}

// NewGameUpdatesPub returns a new updatePub for the provided GameID.
func NewGameUpdatesPub(gameID uuid.UUID) *updatePub[GameUpdate] {
	gameUpdatesBus := GetGameUpdateBus()
	gameUpdatesBus.NewTopic(gameID)

	return &updatePub[GameUpdate]{
		pub:      bus.NewPub(gameUpdatesBus),
		entityID: gameID,
	}
}
