package player

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/store"
)

var PlayerStore *store.IndexedStore[*Player]

func GetPlayerStore() *store.IndexedStore[*Player] {
	if PlayerStore == nil {
		PlayerStore = &store.IndexedStore[*Player]{
			DataMap: make(map[uuid.UUID]*Player),
			Mux: &sync.RWMutex{},
		}
	}
	return PlayerStore
}
