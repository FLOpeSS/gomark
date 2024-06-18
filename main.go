package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Post struct {
	Title string
	Text  string
}

type Posts struct {
	posts []Post
}

const SERVERPORT = ":8000"

func readMdDir(dirName string) []string {
	var posts []string
	dir, err := os.ReadDir(dirName)
	if err != nil {
		log.Printf("Error occurred: %s", err)
	}
	for _, value := range dir {
		name := value.Name()
		fmt.Println(name)
		posts = append(posts, name)
	}
	return posts
}

func readMdFiles(filenames []string) []string {
	var posts []string
	for i, file := range filenames {
		readFiles, err := os.ReadFile("./posts/" + file)
		if err != nil {
			log.Printf("Error while reading file: %s", err)
		}
		readedFiles := strings.TrimSpace(string(readFiles))
		posts = append(posts, string(readedFiles))
		fmt.Printf("Post %d: %s\n", i+1, readFiles)
	}
	return posts
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	dir := readMdDir("./posts")
	result := readMdFiles(dir)
	fmt.Println(result)
	fmt.Println(len(result))
	fmt.Println(result[0])

	file, err := os.ReadFile("./posts/post1.md")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	fmt.Println(string(file))

	post := Post{
		Title: "Testing markdown to html",
		Text:  "Making a lot of texts",
	}

	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("Error while parsing: ", err)
	}

	t.Execute(w, post)
}

func main() {
	http.HandleFunc("/", HomePageHandler)

	err := http.ListenAndServe(SERVERPORT, nil)
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}

}
