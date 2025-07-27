package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id    string `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

var users = []User{}

type Response struct {
	Message string `json:"message"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Api is working now"})
}
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	users = append(users, User{Id: "1", Name: "Vasya", Email: "vasya@mail.ru"})
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	fmt.Println("server staring...")
	http.ListenAndServe(":8080", r)
}
