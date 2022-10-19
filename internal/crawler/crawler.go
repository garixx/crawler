package crawler

import (
	"crawler/internal/cache"
	"crawler/internal/client"
	"crawler/internal/results"
	"errors"
	"fmt"
	"log"
	"sync"
)

type Crawler struct {
	client client.HTTPClient
	cache  *cache.Cache
}

func NewCrawler(client client.HTTPClient, c *cache.Cache) Crawler {
	return Crawler{client: client, cache: c}
}

func (c *Crawler) Crawl(domains []string) (map[string]results.Result, error) {
	length := len(domains)
	if length < 1 {
		return nil, errors.New("urls list is empty")
	}

	resultsMap, err := results.NewResults(domains)
	if err != nil {
		return nil, errors.New("error")
	}

	var wg sync.WaitGroup
	res := results.Results{
		Results: resultsMap,
	}

	for i := 0; i < length; i++ {
		wg.Add(1)

		go func(uri string, wg *sync.WaitGroup) {
			defer wg.Done()
			if cached, ok := c.cache.Get(uri); ok {
				log.Println("Result from cache: ", uri, " ", cached)
				res.Put(uri, cached)
				return
			}
			log.Println("Starting sending request to ", uri)
			resp, err := c.client.Get(uri)
			if err != nil {
				log.Println("error:", err)
				res.Put(uri, results.Result{Uri: uri, Result: err.Error(), FromCache: false})
				c.cache.Put(uri, err.Error())
				return
			}
			log.Println("response: ", uri, " ", resp.StatusCode)
			response := fmt.Sprintf("%d", resp.StatusCode)
			res.Put(uri, results.Result{Uri: uri, Result: response, FromCache: false})
			c.cache.Put(uri, response)
		}(domains[i], &wg)
	}

	wg.Wait()
	return res.Results, nil
}
