package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("Server starting...")
	http.HandleFunc("/", handlerHttp)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handlerHttp(w http.ResponseWriter, r *http.Request) {

	pathSegments := strings.Split(r.URL.Path, "/")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	expectedArc := pathSegments[1]

	jsonInput, err := os.ReadFile("data.json")
	if err != nil {
		logWithServerError(w, "Error reading JSON data file", err)
		return
	}

	var story StoryLine

	err = json.Unmarshal(jsonInput, &story)
	if err != nil {
		logWithServerError(w, "Error parsing JSON data", err)
		return
	}

	t, _ := template.ParseFiles("template.html")

	retVal, ok := story[expectedArc]

	if !ok {
		log.Printf("Path '%s' not found in the data source", expectedArc)
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	err = t.Execute(w, retVal)

	if err != nil {
		logWithServerError(w, "Error parsing template", err)
	}
}

func logWithServerError(w http.ResponseWriter, msg string, err error) {
	log.Printf("%s: %v", msg, err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
