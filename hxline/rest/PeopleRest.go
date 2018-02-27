package rest

import (
	"encoding/json"
	"hxline/model"
	sv "hxline/services"
	"net/http"
)

func GetPeopleEndPoint(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(sv.ReadAllPeople())
}

func CreatePeopleEndPoint(res http.ResponseWriter, req *http.Request) {
	var person model.Person
	json.NewDecoder(req.Body).Decode(&person)
	sv.Create(person)
}
