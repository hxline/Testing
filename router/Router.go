package router

import (
	"log"
	"net/http"

	peopleRest "hxline/rest"

	"github.com/gorilla/mux"
)

func StartRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/people", peopleRest.GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people", peopleRest.CreatePeopleEndPoint).Methods("POST")
	router.HandleFunc("/people", peopleRest.UpdatePeopleEndPoint).Methods("PUT")
	router.HandleFunc("/people", peopleRest.DeletePeopleEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9081", router))
}
