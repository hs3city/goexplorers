package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlerHttp)
	http.ListenAndServe(":8080", nil)
}

func handlerHttp(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(w, "our application, world")
	jsonInput, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var story StoryLine
	err = json.Unmarshal(jsonInput, &story)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("new").ParseFiles("template.html"))

	err = t.Execute(w, story.Intro.Title)

	if err != nil {
		panic(err)
	}

	// w.Write([]byte(string(story.Intro.Title)))

	// func setTemplate(){}
}
