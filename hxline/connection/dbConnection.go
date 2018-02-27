package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Conn() (db *sql.DB) {
	connStr := "user=postgres password=postgres dbname=hxline sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return
}

func CUD(stmt string, args ...interface{}) {
	db := Conn()

	defer func() {
		log.Println("Connection Closed")
		db.Close()
	}()

	statement, error := db.Prepare(stmt)
	checkError(error)

	result, error := statement.Exec(args...)
	checkError(error)

	affectedRow, error := result.RowsAffected()
	log.Println(affectedRow, "rows changed")
	checkError(error)

	log.Println("Success")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
