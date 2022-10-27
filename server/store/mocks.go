package store

type MockStore[T Indexable] struct{}

func (m *MockStore[T]) Lock() {}

func (m *MockStore[T]) Unlock() {}

func (m *MockStore[T]) Store(t T) {}

func (m *MockStore[T]) Load(t *T) error {
	return nil
}

func (m *MockStore[T]) LoadAll(ts *[]T) {}
