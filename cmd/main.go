package main

import (
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
	r := mux.NewRouter()
	r.HandleFunc(walkRoute, handlers.ParseHandler(client)).Methods("POST")

	log.Println("Listening on port: ", *port)
	_ = http.ListenAndServe(":"+strconv.Itoa(*port), r)
}
