package handlers

import (
	"crawler/internal/app"
	"crawler/internal/client"
	"encoding/json"
	"net/http"
)

type urlsToParse struct {
	Urls []string `json:"urls"`
}

func ParseHandler(client client.HTTPClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var urls urlsToParse

		err := json.NewDecoder(r.Body).Decode(&urls)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		got, err := app.Crawl(client, urls.Urls)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := json.Marshal(got)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(output)
	}
}
