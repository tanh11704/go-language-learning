package main

import (
	"fmt"

	"rsc.io/quote"
)

func GetMessage() string {
	return quote.Go()
}

func main() {
	fmt.Println("Hello, Golang Internship 2025!")
	fmt.Println(GetMessage())
}