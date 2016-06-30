package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fatih/color"
	"github.com/lucasmezencio/image-resizer/util"
	"github.com/lucasmezencio/image-resizer/util/array"
)

func main() {
	directory := "."

	flag.StringVar(&directory, "p", directory, "The path to walk in")
	flag.Parse()

	fmt.Printf("Walking recursively through: %s\n\n", directory)

	filepath.Walk(directory, func(path string, fileInfo os.FileInfo, err error) error {
		imageSize := fileInfo.Size()
		twoMegabytes := int64(2e+6)
		maxSize := 2000
		allowedExtensions := []string{"jpg", "jpeg"}

		fileExt := strings.ToLower(filepath.Ext(path))

		if path != "." && !fileInfo.IsDir() && !util.Empty(fileExt) {
			fileExt = fileExt[1:len(fileExt)]

			if imageSize >= twoMegabytes && (array.InArray(fileExt, allowedExtensions) != -1) {
				imageName := fileInfo.Name()

				fmt.Printf("Image %s is larger than 2 MB (%d)\n", imageName, imageSize)

				reader, err := os.Open(path)

				if err != nil {
					fmt.Println("Error opening", color.RedString(imageName))
				}

				img, err := jpeg.Decode(reader)

				if err != nil {
					fmt.Println("Error decoding", color.RedString(imageName))
				}

				bounds := img.Bounds()

				if bounds.Max.X > maxSize {
					X := strconv.Itoa(bounds.Max.X)

					fmt.Printf("Image %s is wider than 2000px (%s)\n\n",
						color.GreenString(imageName),
						color.BlueString(X+"px"))

					destImg := imaging.Resize(img, maxSize, 0, imaging.Lanczos)

					// @TODO: add option to rename file
					fmt.Println(path)

					err := imaging.Save(destImg, path)

					if err != nil {
						panic(err)
					}
				}

				defer reader.Close()
			}
		}

		return nil
	})

	fmt.Println("Ma oeeee!")
}
