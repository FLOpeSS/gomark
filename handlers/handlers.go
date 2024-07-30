package handlers

import (
	"fmt"
	"gomark/readers"
	"html/template"
	"net/http"
	"path"
	"reflect"
)

var dot byte = 46

type Post struct {
	Title []string
	Posts []template.HTML
}

type Posts struct {
	Posts []Post
}

func filterFile(filename string) []byte {
	var filtered []byte
	for i := range filename {
		if filename[i] == dot {
			break
		}
		filtered = append(filtered, filename[i])
	}
	return filtered
}

func TestingFilter(w http.ResponseWriter, r *http.Request) {
	posts := []string{"post1.md", "post2.md"}
	var result1 []string
	for _, value := range posts {
		result1 = append(result1, string(filterFile(value)))
	}
	fmt.Println(result1)
	fmt.Println("Len of filtered result: ", len(result1))
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /request\n")
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)

	// Converting md files to html and appending title and posts to Post structure(inst)
	var inst Post
	for i, value := range files {
		inst.Posts = append(inst.Posts, template.HTML(readers.MdToHTML([]byte(value))))
		inst.Title = append(inst.Title, string(filterFile(dir[i])))
	}

	fmt.Println("Titles: ", inst.Title)
	fmt.Println("Titles, type: ", reflect.TypeOf(inst.Title))

	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("Error while parsing: ", err)
	}

	t.Execute(w, inst)
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
