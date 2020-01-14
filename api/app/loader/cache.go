package loader

import (
	"context"

	"github.com/nicksrandall/dataloader"
	cache "github.com/patrickmn/go-cache"
)

// Cache implements the dataloader.Cache interface
type Cache struct {
	c *cache.Cache
}

// Get gets a value from the cache
func (c *Cache) Get(_ context.Context, key dataloader.Key) (dataloader.Thunk, bool) {
	v, ok := c.c.Get(key.String())
	if ok {
		return v.(dataloader.Thunk), ok
	}
	return nil, ok
}

// Set sets a value in the cache
func (c *Cache) Set(_ context.Context, key dataloader.Key, value dataloader.Thunk) {
	c.c.Set(key.String(), value, 0)
}

// Delete deletes and item in the cache
func (c *Cache) Delete(_ context.Context, key dataloader.Key) bool {
	if _, found := c.c.Get(key.String()); found {
		c.c.Delete(key.String())
		return true
	}
	return false
}

// Clear clears the cache
func (c *Cache) Clear() {
	c.c.Flush()
}
