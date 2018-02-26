package repository

import (
	"fmt"
	connection "hxline/connection"
	"hxline/model"
	"log"

	"github.com/google/uuid"
)

func Insert() {
	uuidGen, error := uuid.NewRandom()
	db := connection.Conn()

	defer func() {
		fmt.Println("Connection Closed")
		db.Close()
	}()

	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println("Inserting data")
		db.Query("INSERT INTO test VALUES($1, 'GOVAL')", uuidGen)
	}
}

func Update(people model.Person) {
	people = model.Person{"11", "22"}
	fmt.Println("Updating data")
}

func Delete(id string) {
	db := connection.Conn()

	defer func() {
		fmt.Println("Connection Closed")
		db.Close()
	}()

	fmt.Println("Deleting data")
	db.Query("DELETE FROM test WHERE id = $1", id)
}

func Get() {
	fmt.Println("Getting data")
	db := connection.Conn()
	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("PING")
		log.Fatal(pingErr)
	} else {
		fmt.Println("CONNECTED")
	}

	db.Close()
}
