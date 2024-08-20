package handlers

import (
	"fmt"
	"gomark/readers"
	"html/template"
	"net/http"
	"path"
	"time"
)

type PostItem struct {
	Title string
	Post  template.HTML
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Printf("GET /request\n")
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)

	// Create a slice of PostItem to hold titles and posts together
	var postItems []PostItem
	for i, value := range files {
		postItems = append(postItems, PostItem{
			Title: string(readers.FilterFileName(dir[i])),
			Post:  template.HTML(readers.MdToHTML([]byte(value))),
		})
	}

	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, postItems) // Pass postItems to the template
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	endTime := time.Since(start)
	fmt.Println("Function time: ", endTime)
}

func TestingTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/js.html")
	if err != nil {
		fmt.Printf("Error while parsing: %s", err)
	}

	t.Execute(w, nil)
}

func TestingHandler(w http.ResponseWriter, r *http.Request) {
	var readed string
	var posts []string
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)

	for _, value := range files {
		readed = readers.MdToHTML([]byte(value))
		posts = append(posts, readed)
	}

	fmt.Println(readed)
	fmt.Println(posts)
	for i := range posts {
		fmt.Printf("Printing one at the time: %s", posts[i])
	}
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/javascript")
	filenames := r.URL.Path[len("/scripts/"):]
	filenames = path.Clean(filenames)
	fmt.Println(filenames)
	http.ServeFile(w, r, "./scripts/"+filenames)
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/css")
	filenames := r.URL.Path[len("/style/"):]
	filenames = path.Clean(filenames)
	fmt.Println(filenames)
	http.ServeFile(w, r, "./style/"+filenames)
}
