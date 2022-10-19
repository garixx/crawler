package results

import (
	"errors"
	"sync"
)

type Result struct {
	Uri       string
	Result    string
	FromCache bool `json:"-"`
}

type Results struct {
	mu      sync.Mutex
	Results map[string]Result
}

func NewResults(domains []string) (map[string]Result, error) {
	if len(domains) < 1 {
		return nil, errors.New("empty URLs list provided")
	}
	results := map[string]Result{}
	for _, v := range domains {
		results[v] = Result{
			Uri:       v,
			Result:    "Not Processed",
			FromCache: false,
		}
	}
	return results, nil
}

func (r *Results) Put(uri string, result Result) {
	r.mu.Lock()
	r.Results[uri] = result
	r.mu.Unlock()
}
