package main

import (
	"fmt"
	"link"
	"log"
	"os"
)

func w() {
	//zrobmy

}
func makeMeTree() {
	//zrobmy

}

func main() {
	file, err := os.Open("./grandehtml.html")
	if err != nil {
		log.Fatalf("Error reading HTML file: %v", err)
	}

	links := link.ParseLinks(*file)

	fmt.Println(links)
}
