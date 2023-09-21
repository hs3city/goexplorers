package main

import "fmt"

func Hello(msg string) string {
	return fmt.Sprintf("Hello %s", msg)
}

func main() {
	fmt.Println(Hello("world"))
}
