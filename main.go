package main

import (
	"fmt"
	"gomark/handlers"
	"html/template"
	"net/http"
)

const SERVERPORT = ":8000"

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/testing", handlers.TestingHandler)

	err := http.ListenAndServe(SERVERPORT, nil)
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}

}
