package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"html/template"
	"path/filepath"
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
	
	fpost := " "
	// Edit template.. 
	tmpl := template.Must(template.ParseFiles("first-post.txt"))
	tmpl.Execute(os.Stdout, fpost)

	const html = "<!doctype html><head><title>Untitled Custom SSG</title></head><body></body></html>"

	// Write HTML template to file 
	t := template.Must(template.New("first-post.html").Parse(html))
	fmt.Print(t)

	// File Flag
	file := flag.String(".txt", ".html", "Update file")
	flag.Parse()
	fmt.Println(*file)

	// Dir Flag 
	// find all the .txt files in the directory & print to standout
	path := "~/go/makesite"
	files, err := ioutil.ReadDir(path)
	
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".txt" {
			fmt.Println(f.Name())
		}
	}

	if err != nil {
		fmt.Print(err)
	}

}