package main

import (
	"database/sql"
	"fmt"
	"log"
	"madorsky_go.site/blog/pkg/utils"
)

var cred *utils.Credentionals

func createTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS articles (" +
		"\n    id SERIAL PRIMARY KEY," +
		"\n    title VARCHAR(255)," +
		"\n    text TEXT\n);" +
		"\n")
	if err != nil {
		log.Fatalln(err)
	}
}

func getConnectionDB() *sql.DB {
	credentials, err := utils.ParseFile("admin_credentials.txt")
	if err != nil {
		log.Fatalln(err)
	}
	cred = credentials

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=golang_db port=5432 sslmode=disable", cred.PgUser, cred.PgPass, cred.PgDB)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	createTable(db)
	return db
}
