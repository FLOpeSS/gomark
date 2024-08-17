package main

import (
	"fmt"
	"gomark/handlers"
	"net/http"
)

const SERVERPORT = ":8000"

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/testing", handlers.TestingHandler)
	http.HandleFunc("/testingTemplate", handlers.TestingTemplate)
	http.HandleFunc("/api/data", handlers.ApiHandler)

	http.HandleFunc("/style/", handlers.CssHandler)
	http.HandleFunc("/scripts/", handlers.JsHandler)

	err := http.ListenAndServe(SERVERPORT, nil)
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}

}
