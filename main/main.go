package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abdulkaderm36/gophercises/html-link-parser/parser"
)

func main() {
	filename := flag.String("f", "", "html file to parse")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

    p := parser.Parser{
        File: file,
    }

    links, err := p.Parse()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v", links)
}
