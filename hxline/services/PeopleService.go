package services

import (
	"hxline/model"
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

func ReadAllPeople() (peoples []model.Person) {
	peopleRows := repo.GetAllPeople()

	for peopleRows.Next() {
		var id string
		var name string
		var age int
		peopleRows.Scan(&id, &name, &age)
		peoples = append(peoples, model.Person{ID: id, Name: name, Age: age})
	}

	return
}
