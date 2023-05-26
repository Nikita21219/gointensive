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
var cred *utils.Credentionals

func main() {
	credentials, err := utils.ParseFile("admin_credentials.txt")
	if err != nil {
		log.Fatalln(err)
	}
	cred = credentials

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/admin", admin).Methods("GET", "POST")
	r.HandleFunc("/login", login).Methods("GET", "POST")
	http.Handle("/", r)

	// Init file static server
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	// Launch web server
	err = http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatalln("Error launch web server:", err)
	}
	_ = db.Close()
}
