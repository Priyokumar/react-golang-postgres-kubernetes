package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() {
	connStr := "postgresql://test:test@localhost/test?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)
}
