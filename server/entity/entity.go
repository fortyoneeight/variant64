package entity

import (
	"github.com/variant64/server/store"
)

// Entity wraps an Indexable struct and manages its access patterns.
type Entity[T store.Indexable] struct {
	store *store.IndexedStore[T]
	Data T
}

// Entity wraps a list of Indexable structs and manages their access patterns.
type EntityList[T store.Indexable] struct {
	store *store.IndexedStore[T]
	Data []T
}

// EntityReadRequest accepts an Entity and performs a read operation on it.
type EntityReadRequest[T store.Indexable] interface {
	Read(*Entity[T]) error
}

// EntityWriteRequest accepts an Entity and performs a write operation on it.
type EntityWriteRequest[T store.Indexable] interface {
	Write(*Entity[T]) error
}

// EntityListReadRequest accepts an EntityList and performs a read operation on it.
type EntityListReadRequest[T store.Indexable] interface {
	Read(*EntityList[T]) error
}

// Load loads the Data field for an Entity.
func (e *Entity[T]) Load() error {
	e.store.Lock()
	defer e.store.Unlock()

	return e.store.Load(&e.Data)
}

// Load loads the Data field for an EntityList.
func (e *EntityList[T]) Load() {
	e.store.Lock()
	defer e.store.Unlock()

	e.store.LoadAll(&e.Data)
}

// Store saves the Entity reference to the store.
func (e *Entity[T]) Store() {
	e.store.Lock()
	defer e.store.Unlock()

	e.store.Store(e.Data)
}
