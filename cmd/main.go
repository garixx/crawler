package main

import (
	"crawler/internal/cache"
	"crawler/internal/crawler"
	"crawler/internal/handlers"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	walkRoute = "/walk"
)

func main() {
	port := flag.Int("port", 8091, "server port")
	flag.Parse()

	client := &http.Client{Timeout: 10 * time.Second}
	empty := make(map[string]string)
	c := crawler.NewCrawler(client, cache.New(empty))

	r := mux.NewRouter()
	r.HandleFunc(walkRoute, handlers.ParseHandler(c)).Methods("POST")

	log.Println("Listening on port: ", *port)
	_ = http.ListenAndServe(":"+strconv.Itoa(*port), r)
}
