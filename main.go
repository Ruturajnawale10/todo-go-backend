package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	s := router.PathPrefix("/api/v1/").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createuser", createUser).Methods("POST")
	s.HandleFunc("/", home).Methods("GET")
	s.HandleFunc("/signin", signInUser).Methods("GET")
	s.HandleFunc("/addtodo", addTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", s)) //Run Server
}
