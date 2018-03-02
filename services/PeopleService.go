package services

import (
	"hxline/model"
	repo "hxline/repositories"
)

func Create(people model.Person) (err error) {
	err = repo.Insert(people)

	return
}

func Update(people model.Person) (result []model.Person, err error) {
	result, err = repo.GetPerson(people.ID)
	if err == nil && len(result) > 0 {
		err = repo.Update(people)
	}

	return
}

func Delete(id string) (result []model.Person, err error) {
	result, err = repo.GetPerson(id)
	if err == nil && len(result) > 0 {
		err = repo.Delete(id)
	}

	return
}

func ReadAllPeople() (peoples []model.Person, err error) {
	peoples, err = repo.GetAllPeople()

	return
}
