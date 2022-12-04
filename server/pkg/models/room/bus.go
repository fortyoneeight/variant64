package room

import (
	"github.com/google/uuid"
	"github.com/variant64/server/pkg/bus"
)

type updatePub[T any] struct {
	pub      *bus.Pub[T]
	entityID uuid.UUID
}

func (u *updatePub[T]) Publish(update T) {
	u.pub.Publish(u.entityID, update)
}

var roomBus *bus.Bus[RoomUpdate]

// GetRoomBus returns the global bus for Game updates.
func GetRoomBus() *bus.Bus[RoomUpdate] {
	if roomBus == nil {
		roomBus = bus.NewBus[RoomUpdate]([]uuid.UUID{})
		roomBus.Start()
	}
	return roomBus
}

// NewRoomsPub returns a new updatePub for the provided GameID.
func NewRoomUpdatePub(roomID uuid.UUID) *updatePub[RoomUpdate] {
	RoomBus := GetRoomBus()
	RoomBus.NewTopic(roomID)

	return &updatePub[RoomUpdate]{
		pub:      bus.NewPub(RoomBus),
		entityID: roomID,
	}
}
