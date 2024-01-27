package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func catch[T any](val T, err error) T {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return val
}

func main() {

	content := catch(os.ReadFile("ex3.html"))

	reader := bytes.NewReader(content)
	tokenizer := html.NewTokenizer(reader)

	// var links []string
	for {
		t := tokenizer.Next()
		if t == html.ErrorToken {
			//fmt.Println("Error token")
			// fmt.Println(t)
			break
		}

		// fmt.Println(z.Token().Attr)
		token := tokenizer.Token()
		if token.Data == "a" && len(token.Attr) > 0 {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
		}

	}

}
