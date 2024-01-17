package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", handlerHttp)
	http.ListenAndServe(":8080", nil)
}

func handlerHttp(w http.ResponseWriter, r *http.Request) {

	pathSegments := strings.Split(r.URL.Path, "/")
	if r.Method != "GET" {
		fmt.Println("Method not allowed")
		return
	}
	expectedArc := pathSegments[1]

	jsonInput, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var story StoryLine

	err = json.Unmarshal(jsonInput, &story)
	if err != nil {
		panic(err)
	}

	t, _ := template.ParseFiles("template.html")

	retVal, ok := story[expectedArc]

	if !ok {
		fmt.Printf("path %s not found in the data source\n", expectedArc)
		return
	}

	err = t.Execute(w, retVal)

	if err != nil {
		fmt.Println("whaa")
		fmt.Println(err)
		panic(err)
	}
}
