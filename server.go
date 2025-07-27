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

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = fmt.Sprintf("%d", len(users)+1)
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "New user appended"})
}

func UpdateUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i, user := range users {
		if user.Id == id {
			json.NewDecoder(r.Body).Decode(&users[i])
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Response{Message: "User data changed"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Message: "No such user"})
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Response{Message: "User deleted"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Message: "User not found"})
}

func main() {
	users = append(users, User{Id: "1", Name: "Vasya", Email: "vasya@mail.ru"})
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", CreateUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUsersHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUsersHandler).Methods("DELETE")
	fmt.Println("server staring...")
	http.ListenAndServe(":8080", r)
}
