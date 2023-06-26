package basicstore

import (
	"fmt"
	"sync"

	"github.com/infinitybotlist/grevolt/cache/store"
)

// Simple implementation of state
//
// This uses a sync.RWMutex to be able to handle concurrent writes
type BasicStore[T any] struct {
	sync.RWMutex

	dataStore map[string]*T

	// Whether or not to track in this state
	Disabled bool
}

// Initialize the state, not needed by Store interface but needed for this specific store
func (s *BasicStore[T]) Init() *BasicStore[T] {
	s.dataStore = make(map[string]*T)
	return s
}

// Print list of key/value pairs, not needed by Store interface
func (s *BasicStore[T]) KeyValuePairs() []string {
	var kvs []string

	for k, v := range s.dataStore {
		kvs = append(kvs, fmt.Sprintln(k, "=>", v))
	}

	return kvs
}

// Is the state usable
func (s *BasicStore[T]) Usable() bool {
	return !s.Disabled
}

// Get an entity from the state
func (s *BasicStore[T]) Get(id string) (*T, error) {
	if s.Disabled {
		return nil, store.ErrDisabled
	}

	if id == "" {
		return nil, store.ErrIdInvalid
	}

	s.RLock()
	defer s.RUnlock()

	if entity, ok := s.dataStore[id]; ok {
		return entity, nil
	}

	return nil, store.ErrNotFound
}

// Set an entity in the state
func (s *BasicStore[T]) Set(id string, entity *T) error {
	if s.Disabled {
		return store.ErrDisabled
	}

	if id == "" {
		return store.ErrIdInvalid
	}

	s.Lock()
	defer s.Unlock()

	if len(s.dataStore) == 0 {
		s.Init()
	}

	s.dataStore[id] = entity

	return nil
}

// Delete an entity from the state
func (s *BasicStore[T]) Delete(id string) error {
	if s.Disabled {
		return store.ErrDisabled
	}

	if id == "" {
		return store.ErrIdInvalid
	}

	s.Lock()
	defer s.Unlock()

	delete(s.dataStore, id)

	return nil
}

// Returns the length of the store
func (s *BasicStore[T]) Length() int {
	return len(s.dataStore)
}
