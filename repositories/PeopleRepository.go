package repositories

import (
	"database/sql"
	"hxline/connection"
	"hxline/model"
	"log"

	"github.com/google/uuid"
)

//Insert into table people
func Insert(people model.Person) (err error) {
	uuidGen, err := uuid.NewRandom()
	checkError(err)

	log.Println("Inserting data")
	err = connection.CUD("INSERT INTO people VALUES($1, $2, $3)", uuidGen, people.Name, people.Age)

	return
}

//Update into table people
func Update(people model.Person) (err error) {
	log.Println("Updating data")
	err = connection.CUD("UPDATE people SET name=$2, age=$3 WHERE id=$1", people.ID, people.Name, people.Age)

	return
}

//Delete into table people
func Delete(id string) (err error) {
	log.Println("Deleting data")
	err = connection.CUD("DELETE FROM people WHERE id = $1", id)

	return
}

//GetAllPeople from table people
func GetAllPeople() (result []model.Person, err error) {
	log.Println("Reading data")
	sql, err := connection.Retrieve("SELECT * FROM people")

	if err == nil {
		result = toArray(sql)
	}

	return
}

//GetPerson from table people
func GetPerson(id string) (result []model.Person, err error) {
	log.Println("Reading data")
	sql, err := connection.Retrieve("SELECT * FROM people WHERE id=$1", id)

	if err == nil {
		result = toArray(sql)
	}

	return
}

func toArray(sql *sql.Rows) (peoples []model.Person) {
	for sql.Next() {
		var id string
		var name string
		var age int
		sql.Scan(&id, &name, &age)
		peoples = append(peoples, model.Person{ID: id, Name: name, Age: age})
	}

	return
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
