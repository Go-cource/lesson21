package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Get request
func createGetRequest() {
	resp, err := http.Get("http://127.0.0.1:8080/users")
	if err != nil {
		fmt.Println("Get error: ", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Sever answer: ", string(body))
}

func main() {
	user := []byte(`{"id":"0", "name": "Dima", "email": "dima@gmail.com"}`)
	resp, err := http.Post("http://127.0.0.1:8080/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		fmt.Println("Request error", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error", err)
		return
	}

	fmt.Println("Server answered: ", string(body))
}
