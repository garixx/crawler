package results

import (
	"errors"
	"sync"
)

type Results struct {
	mu      sync.Mutex
	Results map[string]string
}

func NewResults(domains []string) (map[string]string, error) {
	if len(domains) < 1 {
		return nil, errors.New("empty URLs list provided")
	}
	results := map[string]string{}
	for _, v := range domains {
		results[v] = "Not Processed"
	}
	return results, nil
}

func (r *Results) Put(uri, result string) {
	r.mu.Lock()
	r.Results[uri] = result
	r.mu.Unlock()
}
