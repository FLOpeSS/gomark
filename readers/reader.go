package readers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

func MdToHTML(md []byte) string {
	html := markdown.ToHTML(md, nil, nil)
	return string(html)
}

func ReadMdDir(dirName string) []string {
	var posts []string
	dir, err := os.ReadDir(dirName)
	if err != nil {
		log.Printf("Error occurred: %s", err)
	}
	for i, value := range dir {
		name := value.Name()
		fmt.Println(name)
		if i+1 == len(dir) {
			fmt.Printf("\n")
		}
		posts = append(posts, name)
	}
	return posts
}

func ReadMdFiles(filenames []string) []string {
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
