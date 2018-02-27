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
	log.Fatal(http.ListenAndServe(":9081", router))
}
