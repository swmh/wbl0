package cache

import (
	"context"
	"errors"

	lru "github.com/hashicorp/golang-lru/v2"
)

var ErrNoSuchKey = errors.New("no such key")

type Cache[K comparable, V any] struct {
	cache *lru.Cache[K, V]
}

func New[K comparable, V any](size int) (*Cache[K, V], error) {
	cache, err := lru.New[K, V](size)
	if err != nil {
		return nil, err
	}

	return &Cache[K, V]{
		cache: cache,
	}, nil
}

func (c *Cache[K, V]) Set(_ context.Context, key K, value V) error {
	c.cache.Add(key, value)
	return nil
}

func (c *Cache[K, V]) Get(_ context.Context, key K) (V, error) {
	v, ok := c.cache.Get(key)
	if !ok {
		return v, ErrNoSuchKey
	}
	return v, nil
}

func (c *Cache[K, V]) IsNoSuchKey(err error) bool {
	return errors.Is(err, ErrNoSuchKey)
}
