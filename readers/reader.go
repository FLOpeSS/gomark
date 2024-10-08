package readers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

var dot byte = 46

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
		fmt.Println("unfiltered posts: ", name)
		if i+1 == len(dir) {
			fmt.Printf("\n")
		}
		posts = append(posts, name)
	}
	return posts
}

func ReadMdFiles(filenames []string) []string {
	var posts []string
	for _, file := range filenames {
		readFiles, err := os.ReadFile("./posts/" + file)
		if err != nil {
			log.Printf("Error while reading file: %s", err)
		}
		readedFiles := strings.TrimSpace(string(readFiles))
		posts = append(posts, string(readedFiles))
		// fmt.Printf("Post %d: %s\n", i+1, readFiles)
	}
	return posts
}

func FilterFileName(filename string) []byte {
	var filtered []byte
	for i := range filename {
		if filename[i] == dot {
			break
		}

		filtered = append(filtered, filename[i])
	}

	filtered = bytes.Replace(filtered, []byte("_"), []byte(" "), -1)
	fmt.Printf("Filtered Posts: %s\n", filtered)
	return filtered
}
