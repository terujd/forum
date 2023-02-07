package main

import (
	"fmt"
	"net/http"

	"github.com/terujd/forum/back-end/server/"
)

func main() {
	go server.Start()

	fmt.Println("Forum is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
