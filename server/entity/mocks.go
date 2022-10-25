package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/variant64/server/store"
)

// MockEntity represents a mock Entity for testing.
type MockEntity struct {
	ID    uuid.UUID `json:"id"`
	Field string    `json:"field"`
}

// GetID returns the MockEntity's ID.
func (m MockEntity) GetID() uuid.UUID {
	return m.ID
}

// MockEntityReadRequestSuccess represents a successful ReadRequest to a MockEntity.
type MockEntityReadRequestSuccess struct {
	ID uuid.UUID
}

// Read intializes the ID field of the provided MockEntity.
func (r *MockEntityReadRequestSuccess) Read(e *Entity[MockEntity]) error {
	e.EntityStore = &store.MockStore[MockEntity]{}
	e.Data = MockEntity{
		ID: r.ID,
	}
	return nil
}

// MockEntityReadRequestFailed represents a failed ReadRequest to a MockEntity.
type MockEntityReadRequestFailed struct {
	ID uuid.UUID
}

// Read returns an error representing a failed read.
func (m *MockEntityReadRequestFailed) Read(e *Entity[MockEntity]) error {
	return errors.New("failed to read")
}

// MockEntityListReadRequestSuccess represents a successful ReadRequest to a MockEntityList.
type MockEntityListReadRequestSuccess struct {
	ID uuid.UUID
}

// Read intializes the MockEntityList with one MockEntity.
func (r *MockEntityListReadRequestSuccess) Read(e *EntityList[MockEntity]) error {
	e.EntityStore = &store.MockStore[MockEntity]{}
	e.Data = []MockEntity{
		{ID: r.ID},
	}
	return nil
}

// MockEntityListReadRequestFailed represents a failed ReadRequest to a MockEntityList.
type MockEntityListReadRequestFailed struct {
	ID uuid.UUID
}

// Read returns an error representing a failed read.
func (m *MockEntityListReadRequestFailed) Read(e *EntityList[MockEntity]) error {
	return errors.New("failed to read")
}

// MockEntityWriteRequestSuccess represents a succesful WriteRequest to a MockEntity.
type MockEntityWriteRequestSuccess struct {
	ID    uuid.UUID `json:"id"`
	Field string    `json:"field"`
}

// Write sets the Field field value for the MockEntity.
func (r *MockEntityWriteRequestSuccess) Write(e *Entity[MockEntity]) error {
	e.EntityStore = &store.MockStore[MockEntity]{}
	e.Data = MockEntity{
		ID:    r.ID,
		Field: r.Field,
	}
	return nil
}

// MockEntityWriteRequestSuccess represents a failed WriteRequest to a MockEntity.
type MockEntityWriteRequestFailed struct {
	ID    uuid.UUID `json:"id"`
	Field string    `json:"field"`
}

// Write returns an error representing a failed WriteRequest to a MockEntity.
func (m *MockEntityWriteRequestFailed) Write(e *Entity[MockEntity]) error {
	return errors.New("failed to write")
}
