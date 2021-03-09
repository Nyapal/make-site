package main

import (
	// "fmt"
	"os"
	// "io/ioutil"
	"html/template"
)



type firstpost struct {
	Body string 
}

func main() { 
	fpost := firstpost{" "}
	// fmt.Println("Hello, world!")  
	// Read contents of first-post
	// file, err := ioutil.ReadFile("first-post.txt")
	// if err != nil {
	// 	fmt.Print("No Issues")
	// }
	// fmt.Print(string(file))

	// Edit template.. 
	tmpl := template.Must(template.ParseFiles("first-post.txt"))
	tmpl.Execute(os.Stdout, fpost)
}