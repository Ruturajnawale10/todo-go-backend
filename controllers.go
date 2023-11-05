package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Name     string `bson:"name"`
	UserName string `bson:"username"`
	Password string `bson:"password"`
	Todos    []Todo `bson:"todos"`
}

type Todo struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Completed   bool   `bson:"completed"`
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

func signInUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: signInUser")
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"username": user.UserName}

	var result User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid Credentials")
		return
	}

	fmt.Println("Found a Single Document: ", result)
	json.NewEncoder(w).Encode(result)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home")
}
