package store

import "errors"

// To allow for alternative storage implementations
//
// store provides an interface for caching and storing state
// for a Revolt client
//
// Stores manage the entire read/write mechanism of state *but not the actual caching
// implementation*
type Store[T any] interface {
	// Is the state usable
	Usable() bool

	// Get an entity from the state
	Get(id string) (*T, error)

	// Set an entity in the state
	Set(id string, entity *T) error

	// Delete an entity from the state
	Delete(id string) error

	// Returns the length of the store
	Length() int
}

var ErrNotFound = errors.New("entity not found")
var ErrDisabled = errors.New("state tracking is disabled")
var ErrIdInvalid = errors.New("id is invalid")
