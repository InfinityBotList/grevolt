package resolvers

import (
	"github.com/infinitybotlist/grevolt/rest/restcli"
	"github.com/infinitybotlist/grevolt/types"
)

// Resolves a user
//
// # This tries cache first, then fetches directly
//
// The direct fetch also adds to cache meaning that fetching the same user
// multiple times will lead to at maximum one API call (assuming the same session)
func ResolveUser(r *restcli.RestClient, id string) (*types.User, error) {
	// Check cache
	cachedUser, err := r.Config.SharedState.GetUser(id)

	if err != nil || cachedUser.Username == "" {
		// Not in cache, fetch directly
		user, err := r.FetchUser(id)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return cachedUser, nil

}
