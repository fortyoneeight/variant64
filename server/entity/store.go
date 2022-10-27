package entity

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/store"
)

var playerStore *store.IndexedStore[Player]
var roomStore *store.IndexedStore[Room]
var gameStore *store.IndexedStore[*Game]

// GetPlayerStore returns the global store for Player entities.
func GetPlayerStore() *store.IndexedStore[Player] {
	if playerStore == nil {
		playerStore = &store.IndexedStore[Player]{
			DataMap: make(map[uuid.UUID]Player),
			Mux:     &sync.RWMutex{},
		}
	}
	return playerStore
}

// GetRoomStore returns the global store for Room entities.
func GetRoomStore() *store.IndexedStore[Room] {
	if roomStore == nil {
		roomStore = &store.IndexedStore[Room]{
			DataMap: make(map[uuid.UUID]Room),
			Mux:     &sync.RWMutex{},
		}
	}
	return roomStore
}

// GetGameStore returns the global store for Game entities.
func GetGameStore() *store.IndexedStore[*Game] {
	if gameStore == nil {
		gameStore = &store.IndexedStore[*Game]{
			DataMap: make(map[uuid.UUID]*Game),
			Mux: &sync.RWMutex{},
		}
	}
	return gameStore
}
