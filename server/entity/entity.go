package entity

import (
	"github.com/variant64/server/store"
)

type entityStore[T store.Indexable] interface {
	Lock()
	Unlock()
	Store(t T)
	Load(t *T) error
	LoadAll(ts *[]T)
}

// Entity wraps an Indexable struct and manages its access patterns.
type Entity[T store.Indexable] struct {
	EntityStore entityStore[T]
	Data T
}

// Entity wraps a list of Indexable structs and manages their access patterns.
type EntityList[T store.Indexable] struct {
	EntityStore entityStore[T]
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
	e.EntityStore.Lock()
	defer e.EntityStore.Unlock()

	return e.EntityStore.Load(&e.Data)
}

// Load loads the Data field for an EntityList.
func (e *EntityList[T]) Load() {
	e.EntityStore.Lock()
	defer e.EntityStore.Unlock()

	e.EntityStore.LoadAll(&e.Data)
}

// Store saves the Entity reference to the store.
func (e *Entity[T]) Store() {
	e.EntityStore.Lock()
	defer e.EntityStore.Unlock()

	e.EntityStore.Store(e.Data)
}
