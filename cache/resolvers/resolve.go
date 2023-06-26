package resolvers

import (
	"errors"

	"github.com/infinitybotlist/grevolt/rest/restcli"
	"github.com/infinitybotlist/grevolt/types"
)

type Resolvable[T any] struct {
	string
}

// Returns the ID of the resolvable entity
func (r *Resolvable[T]) Id() string {
	return r.string
}

func (r *Resolvable[T]) Resolve(rc *restcli.RestClient) (*T, error) {
	var t T

	switch any(t).(type) {
	case *types.User:
		e, err := ResolveUser(rc, r.string)
		return any(e).(*T), err
	}

	return nil, errors.New("dont know how to fetch this entity type")
}
