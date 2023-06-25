package cache

import (
	"reflect"

	"github.com/infinitybotlist/grevolt/types"
)

func (c *CachedClient) cacheGeneric(v any) error {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error("Recovered from panic: ", r)
		}
	}()

	if c.logger == nil {
		c.logger = c.Rest.Config.Logger
	}

	c.Lock()
	defer c.Unlock()

	// normalize, all our types right now will be **types.User, indirect it
	value := reflect.ValueOf(v)

	var val any

	rval := reflect.Indirect(value)

	if rval.Type().Kind() == reflect.Ptr {
		val = rval.Interface()
	}

	switch val.(type) {
	case *types.User:
		c.logger.Info("Caching user")
		user := val.(*types.User)

		c.users[user.Id] = user
	}

	return nil
}
