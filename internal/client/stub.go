package client

import (
	"errors"
	"fmt"
	"net/http"
)

/* Test stub HTTP client */
type ResponseMap map[string]func() (resp *http.Response, err error)

type MockClient struct {
	Responses ResponseMap
}

func NewMockClient(responses map[string]func() (resp *http.Response, err error)) MockClient {
	return MockClient{
		Responses: responses,
	}
}

func (c *MockClient) Get(url string) (resp *http.Response, err error) {
	if fn, ok := c.Responses[url]; ok {
		return fn()
	}
	return nil, errors.New(fmt.Sprint("stub processor for url=", url, " not found"))
}
