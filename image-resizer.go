package main

import (
	"flag"

	"github.com/lucasmezencio/image-resizer/resizer"
)

func main() {
	directory := "."
	maxSize := "2Mb"
	maxWidth := "2000"

	flag.StringVar(&directory, "p", directory, "The path to walk in")
	flag.StringVar(&maxSize, "s", maxSize, "The max size of the image file")
	flag.StringVar(&maxWidth, "w", maxWidth, "The max width (in px) of the (re)generated image")
	flag.Parse()

	resizer.DoResize(directory, maxSize, maxWidth)
}
