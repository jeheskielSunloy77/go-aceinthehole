package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func CreateDB() Database {
	db, err := sql.Open("postgres", "postgresql://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return Database{db}
}

func (db *Database) Seed() {
	_, err := db.Exec(`
			INSERT INTO "User" (username, email) VALUES ('user1', 'user1@gmail.com');
			INSERT INTO "User" (username, email) VALUES ('user2', 'user2@gmail.com');
		`)
	if err != nil {
		log.Fatal(err)
	}
}
