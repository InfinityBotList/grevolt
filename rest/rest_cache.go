package rest

import "github.com/infinitybotlist/grevolt/types"

func CacheImpl(r *RequestData, v any) {
	// Doesn't do anything yet
	switch v := v.(type) {
	case *types.User:
		// Add to cache
		r.Config.SharedState.AddUser(v)
	case *types.Emoji:
		// Add to cache
		r.Config.SharedState.AddEmoji(v)
	case *types.Server:
		// Add to cache
		r.Config.SharedState.AddServer(v)
	}
}

func Cacher(r *RequestData, v any) error {
	if r.Config.DisableRestCaching {
		return nil
	}

	go CacheImpl(r, v)
	return nil
}
