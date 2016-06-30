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

func DoResizer(imagePath string, fileInfo os.FileInfo){
	imageName := fileInfo.Name()
	maxSize := 2000

	reader, err := os.Open(imagePath)

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
		fmt.Println(imagePath)

		err := imaging.Save(destImg, imagePath)

		if err != nil {
			panic(err)
		}
	}

	defer reader.Close()
}

// Resizer
func Resizer(directory string) {
	twoMegabytes := int64(2e+6)
	//files := []string{}
	filepath.Walk(directory, func(path string, fileInfo os.FileInfo, err error) error {

		allowedExtensions := []string{"jpg", "jpeg"}

		fileExt := strings.ToLower(filepath.Ext(path))

		if path != "." && !fileInfo.IsDir() && !util.Empty(fileExt) {
			fileExt = fileExt[1:len(fileExt)]
			imageName := fileInfo.Name()
			imageSize := fileInfo.Size()

			if imageSize >= twoMegabytes && (array.InArray(fileExt, allowedExtensions) != -1) {

				fmt.Printf("Image %s is larger than 2 MB (%d)\n", imageName, imageSize)

				//files = append(files, path)
				go DoResizer(path, fileInfo)
			}
		}

		return nil
	})

	fmt.Println("Ma oeeee!")
}
