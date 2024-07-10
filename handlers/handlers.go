package handlers

import (
	"fmt"
	"gomark/readers"
	"html/template"
	"net/http"
	"path"
)

type Post struct {
	Title []string
	Posts []template.HTML
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
	filenames := r.URL.Path[len("/style/"):]
	filenames = path.Clean(filenames)
	fmt.Println(filenames)
	http.ServeFile(w, r, "./style/"+filenames)
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/css")
	filenames := r.URL.Path[len("/style/"):]
	filenames = path.Clean(filenames)
	fmt.Println(filenames)
	http.ServeFile(w, r, "./style/"+filenames)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /request\n")
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)

	// Converting md files to html and appending title and posts to Post structure(inst)
	var inst Post
	for i, value := range files {
		inst.Posts = append(inst.Posts, template.HTML(readers.MdToHTML([]byte(value))))
		inst.Title = append(inst.Title, dir[i])
	}
	fmt.Println("Titles: ", inst.Title)

	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("Error while parsing: ", err)
	}

	t.Execute(w, inst)
}
