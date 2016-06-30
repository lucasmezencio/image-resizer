package main

import (
	"flag"
	"fmt"

	"github.com/lucasmezencio/image-resizer/resizer"
)

func main() {
	directory := "."

	flag.StringVar(&directory, "p", directory, "The path to walk in")
	flag.Parse()

	fmt.Printf("Walking recursively through: %s\n\n", directory)

	resizer.Resizer(directory)
}
