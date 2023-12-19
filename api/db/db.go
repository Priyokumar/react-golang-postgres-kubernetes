package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() {
	log.Println("connecting to db")
	user := os.Getenv("POSTGRES_USER")
	psswd := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	log.Println(user, psswd, database, host)

	connStr := "postgresql://" + user + ":" + psswd + "@" + host + ":5432/" + database + "?sslmode=disable"
	//connStr := "postgresql://postgres:root@localhost:5432/test?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	log.Println(db)
}
