package repositories

import (
	"database/sql"
	"hxline/connection"
	"hxline/model"
	"log"

	"github.com/google/uuid"
)

//Insert into table people
func Insert(people model.Person) {
	uuidGen, error := uuid.NewRandom()
	checkError(error)

	log.Println("Inserting data")
	connection.CUD("INSERT INTO people VALUES($1, $2, $3)", uuidGen, people.Name, people.Age)
}

//Update into table people
func Update(people model.Person) {
	log.Println("Updating data")
	connection.CUD("UPDATE people SET name=$2, age=$3 WHERE id=$1", people.ID, people.Name, people.Age)
}

//Delete into table people
func Delete(id string) {
	log.Println("Deleting data")
	connection.CUD("DELETE FROM people WHERE id = $1", id)
}

//GetAllPeople from table people
func GetAllPeople() (result *sql.Rows) {
	db := connection.Conn()
	defer func() {
		log.Println("Connection Closed")
		db.Close()
	}()

	log.Println("Getting data")
	result, error := db.Query("SELECT * FROM people")
	checkError(error)

	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
