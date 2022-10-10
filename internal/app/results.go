package app

import (
	"errors"
	"sync"
)

type Results struct {
	mu      sync.Mutex
	results map[string]string
}

func NewResultsMap(domains []string) (map[string]string, error) {
	if len(domains) < 1 {
		return nil, errors.New("empty URLs list provided")
	}
	results := map[string]string{}
	for _, v := range domains {
		results[v] = "Not Processed"
	}
	return results, nil
}
