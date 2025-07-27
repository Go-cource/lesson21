package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//Get request
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
