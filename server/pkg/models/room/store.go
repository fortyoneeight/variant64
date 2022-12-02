package room

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/store"
)

var roomStore *store.IndexedStore[*Room]

// getRoomStore returns the global store for Room entities.
func getRoomStore() *store.IndexedStore[*Room] {
	if roomStore == nil {
		roomStore = &store.IndexedStore[*Room]{
			DataMap: make(map[uuid.UUID]*Room),
			Mux:     &sync.RWMutex{},
		}
	}
	return roomStore
}
