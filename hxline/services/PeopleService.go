package services

import (
	"database/sql"
	model "hxline/model"
	repo "hxline/repositories"
)

func Create(people model.Person) {
	repo.Insert(people)
}

func Update(people model.Person) {
	repo.Update(people)
}

func Delete(id string) {
	repo.Delete(id)
}

func ReadAllPeople() *sql.Rows {
	return repo.GetAllPeople()
}
