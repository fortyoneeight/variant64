package game

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/store"
)

var gameStore *store.IndexedStore[*Game]

// getGameStore returns the global store for Game entities.
func getGameStore() *store.IndexedStore[*Game] {
	if gameStore == nil {
		gameStore = &store.IndexedStore[*Game]{
			DataMap: make(map[uuid.UUID]*Game),
			Mux:     &sync.RWMutex{},
		}
	}
	return gameStore
}
