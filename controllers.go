package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Todos    []Todo `json:"todos"`
}

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var userCollection = db().Database("go-todo").Collection("users")

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createUser")
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document of new user: ", insertResult.InsertedID)
	json.NewEncoder(w).Encode(insertResult.InsertedID)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home")
}
