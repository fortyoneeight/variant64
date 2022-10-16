package room

import (
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/store"
)

var RoomStore *store.IndexedStore[*Room]

func GetRoomStore() *store.IndexedStore[*Room] {
	if RoomStore == nil {
		RoomStore = &store.IndexedStore[*Room]{
			DataMap: make(map[uuid.UUID]*Room),
			Mux: &sync.RWMutex{},
		}
	}
	return RoomStore
}
