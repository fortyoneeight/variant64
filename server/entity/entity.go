package entity

import (
	"github.com/variant64/server/store"
)

type Entity[T store.Indexable] struct {
	store *store.IndexedStore[T]
	Data *T
}

type EntityList[T store.Indexable] struct {
	store *store.IndexedStore[T]
	Data *[]*T
}

type EntityReadRequest[T store.Indexable] interface {
	Read(*Entity[T]) error
}

type EntityWriteRequest[T store.Indexable] interface {
	Write(*Entity[T]) error
}

type EntityListReadRequest[T store.Indexable] interface {
	Read(*EntityList[T]) error
}

func (e *Entity[T]) Load() error {
	e.store.Lock()
	defer e.store.Unlock()

	return e.store.Load(e.Data)
}

func (e *EntityList[T]) Load() {
	e.store.Lock()
	defer e.store.Unlock()

	e.store.LoadAll(e.Data)
}

func (e *Entity[T]) Store() {
	e.store.Lock()
	defer e.store.Unlock()

	e.store.Store(*e.Data)
}
