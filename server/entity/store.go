package entity

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/store"
)

var PlayerStore *store.IndexedStore[Player]
var RoomStore *store.IndexedStore[Room]
var GameStore *store.IndexedStore[Game]

func GetPlayerStore() *store.IndexedStore[Player] {
	if PlayerStore == nil {
		PlayerStore = &store.IndexedStore[Player]{
			DataMap: make(map[uuid.UUID]Player),
			Mux:     &sync.RWMutex{},
		}
	}
	return PlayerStore
}

func GetRoomStore() *store.IndexedStore[Room] {
	if RoomStore == nil {
		RoomStore = &store.IndexedStore[Room]{
			DataMap: make(map[uuid.UUID]Room),
			Mux:     &sync.RWMutex{},
		}
	}
	return RoomStore
}

func GetGameStore() *store.IndexedStore[Game] {
	if GameStore == nil {
		GameStore = &store.IndexedStore[Game]{
			DataMap: make(map[uuid.UUID]Game),
			Mux:     &sync.RWMutex{},
		}
	}
	return GameStore
}
