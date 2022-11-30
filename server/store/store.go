package store

import (
	"sync"

	"github.com/google/uuid"
)

type Indexable interface {
	GetID() uuid.UUID
}

type IndexedStore[T Indexable] struct {
	DataMap map[uuid.UUID]T
	Mux     *sync.RWMutex
}

func (i *IndexedStore[T]) Lock() {
	i.Mux.Lock()
}

func (i *IndexedStore[T]) Unlock() {
	i.Mux.Unlock()
}

func (i *IndexedStore[T]) GetAll() []T {
	list := []T{}
	for _, val := range i.DataMap {
		list = append(list, val)
	}
	return list
}

func (i *IndexedStore[T]) GetByID(id uuid.UUID) T {
	var result T
	if val, ok := i.DataMap[id]; ok {
		result = val
	}
	return result
}

func (i *IndexedStore[T]) Store(t T) {
	i.DataMap[t.GetID()] = t
}

func (i *IndexedStore[T]) Delete(id uuid.UUID) {
	delete(i.DataMap, id)
}
