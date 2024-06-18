package main

import (
	"fmt"
	"gomark/readers"
	"net/http"
	"text/template"
)

type Post struct {
	Title string
	Text  string
}

const SERVERPORT = ":8000"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)

	var posted []Post
	for i := range files {
		posted = append(posted, Post{
			Text: files[i],
		})
	}

	fmt.Println(posted[0])
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("Error while parsing: ", err)
	}

	for i := range posted {
		t.Execute(w, posted[i])
	}
}

func main() {
	http.HandleFunc("/", HomePageHandler)

	err := http.ListenAndServe(SERVERPORT, nil)
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}

}
