package rest

import "github.com/infinitybotlist/grevolt/types"

func cacheImpl(r *RequestData, v any) {
	// Doesn't do anything yet
	switch v := v.(type) {
	case *types.User:
		// Add to cache
		r.Config.SharedState.AddUser(v)
	case *types.Emoji:
		// Add to cache
		r.Config.SharedState.AddEmoji(v)
	}
}

func cache(r *RequestData, v any) error {
	if r.Config.DisableRestCaching {
		return nil
	}

	go cacheImpl(r, v)
	return nil
}
