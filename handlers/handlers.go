package handlers

import (
	"fmt"
	"gomark/readers"
	"html/template"
	"net/http"
	"path"
)

type Posts struct {
	Posts []Post
}

type Post struct {
	Title    string
	Contents []template.HTML
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /request\n")
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)

	// Converting md files to html and appending title and posts to Post structure(inst)
	var inst Post
	var Posters Posts
	for i, value := range files {
		inst.Contents = append(inst.Contents, template.HTML(readers.MdToHTML([]byte(value))))
		// inst.Contents = template.HTML(readers.MdToHTML([]byte(value)))
		// inst.Contents = template.HTML(readers.MdToHTML([]byte(value)))
		inst.Title = dir[i]
		Posters.Posts = append(Posters.Posts, inst)
	}

	fmt.Printf("Posters: %d\n", len(Posters.Posts))
	fmt.Printf("Posters: %s\n", Posters)

	// for i, value := range Posters.Posts {
	// 	fmt.Printf("Post %d: %s", i+1, value)
	// }

	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("Error while parsing: ", err)
	}

	t.Execute(w, Posters)
}

func TestingHandler(w http.ResponseWriter, r *http.Request) {
	var readed string
	var reading []string
	var posts []string
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)

	for _, value := range files {
		readed = readers.MdToHTML([]byte(value))
		reading = append(reading, value)
		posts = append(posts, readed)
	}

	fmt.Println("readed ", readed)
	fmt.Println("reading: ", reading)
	// fmt.Println(posts)
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
