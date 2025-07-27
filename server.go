package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Api is working now"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	fmt.Println("server staring...")
	http.ListenAndServe(":8080", r)
}
