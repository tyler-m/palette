package palette

import (
	"testing"
)

func TestCreate(t *testing.T) {
	type CreateTest struct {
		name     string
		images   []string
		k        int
		seed     int64
		dsFactor float64
		clusters []Cluster
	}

	testdataImagesPath := "../testdata/images/"
	createTests := make([]CreateTest, 0)

	aerialTest := CreateTest{
		name:     "aerial create test",
		images:   []string{testdataImagesPath + "aerial.jpg"},
		k:        5,
		seed:     21539346,
		dsFactor: 1,
		clusters: []Cluster{
			{R: 135, G: 141, B: 142, A: 255},
			{R: 173, G: 174, B: 180, A: 255},
			{R: 14, G: 168, B: 177, A: 255},
			{R: 97, G: 109, B: 119, A: 255},
			{R: 201, G: 166, B: 103, A: 255}}}

	createTests = append(createTests, aerialTest)

	for _, test := range createTests {
		t.Run(test.name, func(t *testing.T) {
			output := Create(test.images, test.k, test.seed, test.dsFactor)

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
