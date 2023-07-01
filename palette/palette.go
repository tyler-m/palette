package palette

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
)

type Pixel struct {
	x, y int
}

type Color struct {
	color  color.RGBA
	pixels []Pixel
}

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

func getColors(image image.Image) []Color {
	width, height := image.Bounds().Max.X, image.Bounds().Max.Y
	colorToPixelsMap := make(map[color.RGBA][]Pixel)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixel := Pixel{x: x, y: y}
			color := color.RGBAModel.Convert(image.At(x, y)).(color.RGBA)
			colorToPixelsMap[color] = append(colorToPixelsMap[color], pixel)
		}
	}

	colors := make([]Color, 0)
	for color, pixels := range colorToPixelsMap {
		colors = append(colors, Color{color: color, pixels: pixels})
	}

	return colors
}

func Create(imagePaths []string, k int) string {
	for _, imagePath := range imagePaths {
		image := loadImage(imagePath)
		getColors(image)
	}

	return "Well met!"
}
