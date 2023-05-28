package main

import (
	_ "database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	_ "madorsky_go.site/blog/pkg/models"
	"madorsky_go.site/blog/pkg/utils"
	_ "madorsky_go.site/blog/pkg/utils"
	"net/http"
)

var db = getConnectionDB()

func main() {
	// Routes
	r := mux.NewRouter()
	r.Path("/").Methods("GET").Handler(utils.RateLimiter(index))
	r.Path("/admin").Methods("GET", "POST").Handler(utils.RateLimiter(admin))
	r.Path("/login").Methods("GET", "POST").Handler(utils.RateLimiter(login))
	r.Path("/article/{id:[0-9]+}").Methods("GET").Handler(utils.RateLimiter(article))
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
