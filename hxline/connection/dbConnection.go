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

func CUD(stmt string, params ...interface{}) (err error) {
	db := Conn()

	defer func() {
		log.Println("Connection Closed")
		db.Close()
	}()

	statement, err := db.Prepare(stmt)
	checkError(err)

	if err != nil {
		return
	}

	result, err := statement.Exec(params...)
	checkError(err)

	if err != nil {
		return
	}

	affectedRow, err := result.RowsAffected()
	log.Println(affectedRow, "rows changed")
	checkError(err)

	log.Println("Success")

	return
}

func Retrieve(stmt string, params ...interface{}) (result *sql.Rows, err error) {
	db := Conn()

	defer func() {
		log.Println("Connection Closed")
		db.Close()
	}()

	result, err = db.Query(stmt, params...)
	checkError(err)

	log.Println("Success")

	return
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
