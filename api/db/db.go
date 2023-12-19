package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() {
	log.Println("connecting to db")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	if pgUser == "" {
		pgUser = "postgres"
	}
	if pgPassword == "" {
		pgPassword = "root"
	}
	if database == "" {
		database = "test"
	}
	if host == "" {
		host = "localhost"
	}

	log.Println(pgUser, pgPassword, database, host)

	connStr := "postgresql://" + pgUser + ":" + pgPassword + "@" + host + ":5432/" + database + "?sslmode=disable"
	//connStr := "postgresql://postgres:root@localhost:5432/test?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")
	log.Println("pinging...")
	err = db.Ping()
	if err != nil {
		log.Fatal(" DB ping error: ", err.Error())
	}
	log.Println("pong...")

}
