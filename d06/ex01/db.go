package main

import (
	"database/sql"
	"log"
)

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
	// TODO fix access connection
	connStr := "user=postgres password=postgres dbname=blog port=5444 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	createTable(db)
	return db
}
