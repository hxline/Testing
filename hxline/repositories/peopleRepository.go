package repositories

import (
	connection "hxline/connection"
	"hxline/model"
	"log"

	"github.com/google/uuid"
)

func Insert(people model.Person) {
	uuidGen, error := uuid.NewRandom()
	db := connection.Conn()

	defer func() {
		log.Println("Connection Closed")
		db.Close()
	}()

	if error != nil {
		log.Println(error)
	} else {
		log.Println("Inserting data")
		rows, error := db.Query("INSERT INTO people VALUES($1, $2, $3)", uuidGen, people.Name, people.Age)
		if error != nil {
			log.Println(rows)
			log.Println(error)
		} else {
			log.Println(rows)
			log.Println("SUCCESS")
		}
	}
}

func Update(people model.Person) {
	db := connection.Conn()

	defer func() {
		log.Println("Connection Closed")
		db.Close()
	}()

	log.Println("Updating data")
	rows, error := db.Query("UPDATE people SET name=$2, age=$3 WHERE id=$1", people.Id, people.Name, people.Age)
	if error != nil {
		log.Println(rows)
		log.Println(error)
	} else {
		log.Println(rows)
		log.Println("SUCCESS")
	}
}

func Delete(id string) {
	db := connection.Conn()

	defer func() {
		log.Println("Connection Closed")
		db.Close()
	}()

	log.Println("Deleting data")
	rows, error := db.Query("DELETE FROM people WHERE id = $1", id)
	if error != nil {
		log.Println(rows)
		log.Println(error.Error())
	} else {
		log.Println(rows)
		log.Println("SUCCESS")
	}
}

func Get() {
	log.Println("Getting data")
	db := connection.Conn()
	pingErr := db.Ping()

	if pingErr != nil {
		log.Println("PING")
		log.Println(pingErr)
	} else {
		log.Println("CONNECTED")
	}

	db.Close()
}
