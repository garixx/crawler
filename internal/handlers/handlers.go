package handlers

import (
	"crawler/internal/crawler"
	"encoding/json"
	"net/http"
)

type task struct {
	Urls []string `json:"urls"`
}

func ParseHandler(a crawler.Crawler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var task task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		got, err := a.Crawl(task.Urls)
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
