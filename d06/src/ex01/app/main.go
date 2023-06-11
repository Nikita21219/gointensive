package main

import (
	_ "database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	_ "madorsky_go.site/blog/pkg/models"
	"net/http"
)

var db = getConnectionDB()

func main() {
	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/admin", admin).Methods("GET", "POST")
	r.HandleFunc("/login", login).Methods("GET", "POST")
	r.HandleFunc("/article/{id:[0-9]+}", article).Methods("GET")
	http.Handle("/", r)

	// Init file static server
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	// Launch web server
	err := http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatalln("Error launch web server:", err)
	}
	_ = db.Close()
}
