package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type AddTodoRequest struct {
	UserName    string `bson:"username"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
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

	user.Todos = []Todo{}
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

func addTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addTodo")
	w.Header().Set("Content-type", "application/json")

	var addTodo AddTodoRequest
	err := json.NewDecoder(r.Body).Decode((&addTodo))

	fmt.Println(addTodo)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"username": addTodo.UserName}
	newTodo := Todo{Title: addTodo.Title, Description: addTodo.Description, Completed: false}
	update := bson.M{"$push": bson.M{"todos": newTodo}}

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(newTodo)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getTodo")
	w.Header().Set("Content-type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"username": user.UserName}
	opts := options.FindOne().SetProjection(bson.D{{Key: "todos", Value: 1}})
	var result User
	err = userCollection.FindOne(context.TODO(), filter, opts).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	inCompleteTodos := []Todo{}

	for _, todo := range result.Todos {
		if todo.Completed == false {
			inCompleteTodos = append(inCompleteTodos, todo)
		}
	}
	json.NewEncoder(w).Encode(inCompleteTodos)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home")
}
