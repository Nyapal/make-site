package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"html/template"
)

type firstpost struct {
	Body string 
}

func main() { 
	// fmt.Println("Hello, world!") 

	// Read contents of first-post
	// file, err := ioutil.ReadFile("first-post.txt")
	// if err != nil {
	// 	fmt.Print("Houston, we have a prolem")
	// }
	// fmt.Print(string(file))

	// Edit template.. 
	// tmpl := template.Must(template.ParseFiles("first-post.txt"))
	// tmpl.Execute(os.Stdout, fpost)

	// const html = "<!doctype html><head><title>Untitled Custom SSG</title></head><body></body></html>"

	// Write HTML template to file 
	t := template.Must(template.New("first-post.html").Parse(html))
	fmt.Print(t)

	// File Flag... 
	file := flag.String(".txt", ".html", "Update file")
	flag.Parse()
	fmt.Println(*file)

	// Dir Flag 
	dir := flag.String()
	
	// find all the .txt files in the directory & print to standout
}