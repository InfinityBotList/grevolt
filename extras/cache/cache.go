// Package cache provides a simple cache for grevolt
package cache

import (
	"sync"

	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/types"
	"go.uber.org/zap"
)

// CachedClient provides a simple cache for grevolt
type CachedClient struct {
	sync.RWMutex
	client.Client

	// Logger to use, will be autofilled if not provided
	logger *zap.SugaredLogger

	// Users in the cache
	users map[string]*types.User
}

func New() *CachedClient {
	cli := client.New()

	cachedClient := CachedClient{
		Client: cli,
		users:  make(map[string]*types.User),
	}

	/*cli.Rest.Config.OnSuccessfulParse = append(cli.Rest.Config.OnSuccessfulParse, func(r rest.ClientRequest, resp *rest.ClientResponse, v any) error {
		go cachedClient.cacheGeneric(v) // This may be slow, so goroutine it
		return nil
	})*/

	return &cachedClient
}
