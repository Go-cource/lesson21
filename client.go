package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
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

func createPostRequest() {
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

func createPutRequest(id int) {
	client := &http.Client{
		Timeout: 4 * time.Second,
	}
	user := []byte(fmt.Sprintf(`{"id":"%d", "name": "Dmitry", "email": "Dmitry@gmail.com"}`, id))
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://127.0.0.1:8080/users/%d", id), bytes.NewBuffer(user))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "HelloFromCookie")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
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

func main() {
	id := 4
	client := &http.Client{
		Timeout: 4 * time.Second,
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:8080/users/%d", id), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
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
