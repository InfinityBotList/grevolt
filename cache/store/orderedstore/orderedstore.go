package orderedstore

import (
	"sync"

	"github.com/infinitybotlist/grevolt/cache/store"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// An orderedstore is a basic store that preserves insert order
// by using an orderedmap instead of a map
type OrderedStore[T any] struct {
	sync.RWMutex

	dataStore *orderedmap.OrderedMap[string, *T]

	// Whether or not to track in this state
	Disabled bool
}

// Initialize the state
func (s *OrderedStore[T]) Init() *OrderedStore[T] {
	s.dataStore = orderedmap.New[string, *T]()
	return s
}

// Is the state usable
func (s *OrderedStore[T]) Usable() bool {
	return !s.Disabled
}

func (s *OrderedStore[T]) Get(id string) (*T, error) {
	if s.Disabled {
		return nil, store.ErrDisabled
	}

	if id == "" {
		return nil, store.ErrIdInvalid
	}

	s.RLock()
	defer s.RUnlock()

	if entity, ok := s.dataStore.Get(id); ok {
		return entity, nil
	}

	return nil, store.ErrNotFound
}

// Set an entity in the state
func (s *OrderedStore[T]) Set(id string, entity *T) error {
	if s.Disabled {
		return store.ErrDisabled
	}

	if id == "" {
		return store.ErrIdInvalid
	}

	s.Lock()
	defer s.Unlock()

	if s.dataStore.Len() == 0 {
		s.Init()
	}

	s.dataStore.Set(id, entity)

	return nil
}

// Delete an entity from the state
func (s *OrderedStore[T]) Delete(id string) error {
	if s.Disabled {
		return store.ErrDisabled
	}

	if id == "" {
		return store.ErrIdInvalid
	}

	s.Lock()
	defer s.Unlock()

	s.dataStore.Delete(id)

	return nil
}

// Returns the length of the store
func (s *OrderedStore[T]) Length() int {
	return s.dataStore.Len()
}
