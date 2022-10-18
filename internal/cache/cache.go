package cache

import "sync"

type Cache struct {
	mu      sync.Mutex
	results map[string]string
}

func New(results map[string]string) *Cache {
	return &Cache{
		results: results,
	}
}

func (c *Cache) Put(uri, result string) {
	c.mu.Lock()
	c.results[uri] = result
	c.mu.Unlock()
}

func (c *Cache) Get(uri string) (res string, ok bool) {
	res, ok = c.results[uri]
	return
}
