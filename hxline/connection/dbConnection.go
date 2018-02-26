package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Conn() *sql.DB {
	connStr := "user=postgres password=postgres dbname=hxline sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
