package rest

import (
	"encoding/json"
	"hxline/model"
	sv "hxline/services"
	"net/http"
)

func GetPeopleEndPoint(res http.ResponseWriter, req *http.Request) {
	result, err := sv.ReadAllPeople()
	if err != nil {
		res.WriteHeader(http.StatusNoContent)
		json.NewEncoder(res).Encode(model.Person{})
	} else {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(result)
	}
}

func CreatePeopleEndPoint(res http.ResponseWriter, req *http.Request) {
	var person model.Person
	json.NewDecoder(req.Body).Decode(&person)
	err := sv.Create(person)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	} else {
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode("")
	}
}

func UpdatePeopleEndPoint(res http.ResponseWriter, req *http.Request) {
	var person model.Person
	json.NewDecoder(req.Body).Decode(&person)
	people, err := sv.Update(person)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	} else if len(people) < 1 {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("ID not found")
	} else {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode("")
	}
}

func DeletePeopleEndPoint(res http.ResponseWriter, req *http.Request) {
	var person model.Person
	json.NewDecoder(req.Body).Decode(&person)
	people, err := sv.Delete(person.ID)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	} else if len(people) < 1 {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("ID not found")
	} else {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode("")
	}
}
