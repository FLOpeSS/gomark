package handlers

import (
	"fmt"
	"gomark/readers"
	"html/template"
	"net/http"
	"path"
)

type Post struct {
	Title string
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

//
// func JsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/javascript")
// 	filenames := r.URL.Path[len("/style/"):]
// 	filenames = path.Clean(filenames)
// 	fmt.Println(filenames)
// 	http.ServeFile(w, r, "./style/"+filenames)
// }

func CssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/css")
	filenames := r.URL.Path[len("/style/"):]
	filenames = path.Clean(filenames)
	fmt.Println(filenames)
	http.ServeFile(w, r, "./style/"+filenames)
}

// func CssHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "text/css")
//
// 	filename := r.URL.Path[len("/fonts/style/"):]
// 	filename = path.Clean(filename)
//
// 	http.ServeFile(w, r, "./fonts/style/"+filename)
// }

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	dir := readers.ReadMdDir("./posts")
	files := readers.ReadMdFiles(dir)
	var mds []string
	for _, value := range files {
		mds = append(mds, readers.MdToHTML([]byte(value)))
	}

	var inst Post
	for i := range mds {
		inst.Posts = append(inst.Posts, template.HTML(mds[i]))
	}

	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println("Error while parsing: ", err)
	}

	for i := range inst.Posts {
		fmt.Println(inst.Posts[i])
	}
	t.Execute(w, inst)

}
