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

	http.HandleFunc("/style/", handlers.CssHandler)

	err := http.ListenAndServe(SERVERPORT, nil)
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}

}
