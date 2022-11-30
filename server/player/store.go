package player

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/store"
)

var playerStore *store.IndexedStore[*Player]

// getPlayerStore returns the global store for Player entities.
func getPlayerStore() *store.IndexedStore[*Player] {
	if playerStore == nil {
		playerStore = &store.IndexedStore[*Player]{
			DataMap: make(map[uuid.UUID]*Player),
			Mux:     &sync.RWMutex{},
		}
	}
	return playerStore
}
