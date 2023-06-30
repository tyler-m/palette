package palette

import (
	"image"
	"os"
)

func loadImage(imagePath string) image.Image {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return nil
	}

	return image
}

func Create(imagePaths []string, k int) string {
	for _, imagePath := range imagePaths {
		loadImage(imagePath)
	}

	return "Well met!"
}
