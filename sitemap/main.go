package main

import (
	"fmt"
	"link"
	"log"
	"os"
)

type Data struct {
	Lp   int
	Body string
}

type Tree struct {
	Lp   int
	Body string
}

func w(path string) (*os.File, error) {
	file, err := os.Open(path)
	return file, err

}
func makeMeTree(data Data) Tree {
	return Tree{}
}

func main() {
	file, err := w("./grandehtml.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	links := link.ParseLinks(*file)

	fmt.Println(links)
}
