package crawler

import (
	"crawler/internal/cache"
	"crawler/internal/client"
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestSiteCrawler_Crawl(t *testing.T) {
	tests := []struct {
		name      string
		responses client.ResponseMap
		domains   []string
		want      map[string]string
		wantErr   bool
	}{
		{
			name:    "empty url list will fail",
			domains: []string{},
			wantErr: true,
		},
		{
			name:    "one url in list",
			domains: []string{"client://aa.com"},
			responses: client.ResponseMap{
				"client://aa.com": func() (resp *http.Response, err error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			want: map[string]string{
				"client://aa.com": "200",
			},
		},
		{
			name:    "many url with error",
			domains: []string{"client://gg.com", "muu.com", "client://mn.com"},
			responses: client.ResponseMap{
				"client://gg.com": func() (resp *http.Response, err error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
				"muu.com": func() (resp *http.Response, err error) {
					return nil, errors.New("imitate client error")
				},
				"client://mn.com": func() (resp *http.Response, err error) {
					return &http.Response{StatusCode: http.StatusUnauthorized}, nil
				},
			},
			want: map[string]string{
				"client://gg.com": "200",
				"muu.com":         "imitate client error",
				"client://mn.com": "401",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &client.MockClient{Responses: tt.responses}
			empty := make(map[string]string)
			crawler := NewCrawler(mock, cache.New(empty))

			got, err := crawler.Crawl(tt.domains)
			if (err != nil) != tt.wantErr {
				t.Errorf("Crawl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("Crawl() different length: got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Crawl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
