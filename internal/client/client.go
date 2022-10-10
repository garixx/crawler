package client

import (
	"net/http"
)

// HTTPClient Requests executor
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}
