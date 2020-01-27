package main

import (
	"net/http"

	"github.com/ibadsiddiqui/APIs-For-Shahrukh/src/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
