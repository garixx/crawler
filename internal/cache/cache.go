package cache

import (
	"crawler/internal/results"
	"sync"
)

type Cache struct {
	mu      sync.Mutex
	results map[string]results.Result
}

func New(results map[string]results.Result) *Cache {
	return &Cache{
		results: results,
	}
}

func (c *Cache) Put(uri, result string) {
	c.mu.Lock()
	c.results[uri] = results.Result{
		Uri:       uri,
		Result:    result,
		FromCache: true,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(uri string) (res results.Result, ok bool) {
	res, ok = c.results[uri]
	return
}
