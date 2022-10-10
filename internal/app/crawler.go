package app

import (
	"crawler/internal/client"
	"errors"
	"fmt"
	"log"
	"sync"
)

func Crawl(client client.HTTPClient, domains []string) (map[string]string, error) {
	length := len(domains)
	if length < 1 {
		return nil, errors.New("urls list is empty")
	}

	resultsMap, err := NewResultsMap(domains)
	if err != nil {
		return nil, errors.New("error")
	}

	var wg sync.WaitGroup
	res := Results{
		results: resultsMap,
	}

	for i := 0; i < length; i++ {
		wg.Add(1)

		go func(uri string, wg *sync.WaitGroup) {
			defer wg.Done()
			log.Println("Starting sending request to ", uri)
			resp, err := client.Get(uri)
			if err != nil {
				log.Println("error:", err)
				res.mu.Lock()
				res.results[uri] = err.Error()
				res.mu.Unlock()
				return
			}
			log.Println("response: ", uri, " ", resp.StatusCode)
			res.mu.Lock()
			res.results[uri] = fmt.Sprintf("%d", resp.StatusCode)
			res.mu.Unlock()
		}(domains[i], &wg)
	}

	wg.Wait()
	return res.results, nil
}
