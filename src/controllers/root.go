package controllers

import (
	"net/http"

	"github.com/ibadsiddiqui/APIs-For-Shahrukh/src/handlers"
)

func executeAllControllers() {
	http.HandleFunc("/", handlers.IndexHandler)

}
