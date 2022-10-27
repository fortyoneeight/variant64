package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/variant64/server/store"
)

type MockEntity struct {
	ID    uuid.UUID `json:"id"`
	Field string    `json:"field"`
}

func (m MockEntity) GetID() uuid.UUID {
	return m.ID
}

type MockEntityReadRequestSuccess struct {
	ID uuid.UUID
}

func (r *MockEntityReadRequestSuccess) Read(e *Entity[MockEntity]) error {
	e.EntityStore = &store.MockStore[MockEntity]{}
	e.Data = MockEntity{
		ID: r.ID,
	}
	return nil
}

type MockEntityReadRequestFailed struct {
	ID uuid.UUID
}

func (m *MockEntityReadRequestFailed) Read(e *Entity[MockEntity]) error {
	return errors.New("failed to read")
}

type MockEntityListReadRequestSuccess struct {
	ID uuid.UUID
}

func (r *MockEntityListReadRequestSuccess) Read(e *EntityList[MockEntity]) error {
	e.EntityStore = &store.MockStore[MockEntity]{}
	e.Data = []MockEntity{
		{ID: r.ID},
	}
	return nil
}

type MockEntityListReadRequestFailed struct {
	ID uuid.UUID
}

func (m *MockEntityListReadRequestFailed) Read(e *EntityList[MockEntity]) error {
	return errors.New("failed to read")
}

type MockEntityWriteRequestSuccess struct {
	ID    uuid.UUID `json:"id"`
	Field string    `json:"field"`
}

func (r *MockEntityWriteRequestSuccess) Write(e *Entity[MockEntity]) error {
	e.EntityStore = &store.MockStore[MockEntity]{}
	e.Data = MockEntity{
		ID:    r.ID,
		Field: r.Field,
	}
	return nil
}

type MockEntityWriteRequestFailed struct {
	ID    uuid.UUID `json:"id"`
	Field string    `json:"field"`
}

func (m *MockEntityWriteRequestFailed) Write(e *Entity[MockEntity]) error {
	return errors.New("failed to write")
}
