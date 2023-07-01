package palette

import (
	"image/color"
	"testing"
)

func TestCreate(t *testing.T) {
	type CreateTest struct {
		name     string
		images   []string
		k        int
		seed     int64
		clusters []Cluster
	}

	testdataImagesPath := "../testdata/images/"
	createTests := make([]CreateTest, 0)

	aerialTest := CreateTest{
		name:   "aerial create test",
		images: []string{testdataImagesPath + "aerial.jpg"},
		k:      5,
		seed:   21539346,
		clusters: []Cluster{
			{meanColor: color.RGBA{R: 131, G: 140, B: 144, A: 255}},
			{meanColor: color.RGBA{R: 174, G: 175, B: 180, A: 255}},
			{meanColor: color.RGBA{R: 12, G: 168, B: 178, A: 255}},
			{meanColor: color.RGBA{R: 95, G: 107, B: 116, A: 255}},
			{meanColor: color.RGBA{R: 201, G: 166, B: 103, A: 255}}}}

	createTests = append(createTests, aerialTest)

	for _, test := range createTests {
		t.Run(test.name, func(t *testing.T) {
			output := Create(test.images, test.k, test.seed)

			var expectedOutput string
			for _, imagePath := range test.images {
				expectedOutput += Format(imagePath, test.clusters)
			}

			if output != expectedOutput {
				t.Errorf("received output:\n %s \n expected output:\n %s", output, expectedOutput)
			}
		})
	}
}
