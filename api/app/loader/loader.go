package loader

import (
	"context"
	"fmt"
	"time"

	"github.com/nicksrandall/dataloader"
	cache "github.com/patrickmn/go-cache"
)

type key string

const (
	userLoaderKey key = "user"
)

// Client ...
type Client interface {
	userGetter
}

// Initialize ...
func Initialize(client Client) Collection {
	return Collection{
		lookup: map[key]dataloader.BatchFunc{
			userLoaderKey: newUserLoader(client),
		},
	}
}

// Collection ...
type Collection struct {
	lookup map[key]dataloader.BatchFunc
}

// Attach ...
func (c Collection) Attach(ctx context.Context) context.Context {
	for k, batchFn := range c.lookup {
		c := cache.New(15*time.Minute, 15*time.Minute)
		cache := &Cache{c}
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFn, dataloader.WithCache(cache)))
	}

	return ctx
}

func extract(ctx context.Context, k key) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", k)
	}

	return ldr, nil
}
