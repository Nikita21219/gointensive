package main

import (
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Create ES client
var es = getClientConnection()

func getClientConnection() *elastic.Client {
	// Create http client
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
		},
	}

	// Create ES client
	es, err := elastic.NewClient(elastic.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: httpClient.Transport,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	}
	return es
}

func main() {

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	http.Handle("/", r)

	// Launch web server
	err := http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatalln("Error launch web server:", err)
	}
}
