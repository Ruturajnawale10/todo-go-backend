package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	s := router.PathPrefix("/api/v1").Subrouter() //Base Path

	log.Fatal(http.ListenAndServe(":8080", s)) //Run Server
}
