package resizer

import (
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

// DoResize into a directory
func DoResize(directory string, size string, maxWidth string) {
	fmt.Printf("Walking recursively through: %s\n\n", directory)

	bytes := util.StrToBytes(size)

	filepath.Walk(directory, func(path string, fileInfo os.FileInfo, err error) error {
		imageSize := float64(fileInfo.Size())
		allowedExtensions := []string{"jpg", "jpeg"}
		fileExt := strings.ToLower(filepath.Ext(path))
		width, _ := strconv.Atoi(maxWidth)

		if path != "." && !fileInfo.IsDir() && !util.Empty(fileExt) {
			fileExt = fileExt[1:len(fileExt)]

			if imageSize >= bytes && (array.InArray(fileExt, allowedExtensions) != -1) {
				imageName := fileInfo.Name()

				fmt.Printf("Image %s is larger than 2 MB (%f)\n", imageName, imageSize)

				reader, err := os.Open(path)

				if err != nil {
					fmt.Println("Error opening", color.RedString(imageName))
				}

				img, err := jpeg.Decode(reader)

				if err != nil {
					fmt.Println("Error decoding", color.RedString(imageName))
				}

				bounds := img.Bounds()

				if bounds.Max.X > width {
					imageWidth := strconv.Itoa(bounds.Max.X)

					fmt.Printf("Image %s is wider than 2000px (%s)\n\n",
						color.GreenString(imageName),
						color.BlueString(imageWidth+"px"))

					// @TODO: add flag to choose which filter to use
					destImg := imaging.Resize(img, width, 0, imaging.Blackman)

					// @TODO: add flag to add a prefix to filename
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
