package main

import (
	"fmt"
	// "os"
	// "io"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		fmt.Print("No Issues")
	}
	fmt.Print(string(file))
	// fmt.Println("Hello, world!") // Print some ish 
}