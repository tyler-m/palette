package palette

import (
	"image/color"
	"math"
	"math/rand"
)

type Cluster struct {
	meanColor color.RGBA
}

func distance(p Color, q Color) float64 {
	rDifference := q.color.R - p.color.R
	bDifference := q.color.B - p.color.B
	gDifference := q.color.G - p.color.G
	sumOfSquares := rDifference*rDifference + bDifference*bDifference + gDifference*gDifference
	return math.Sqrt(float64(sumOfSquares))
}

func initClusters(colors []Color, k int) []Cluster {
	rng := rand.New(rand.NewSource(8472683430)) // temp for debugging
	clusters := make([]Cluster, k)

	for i, cluster := range clusters {
		colorIndex := int(float64(len(colors)) * rng.Float64())
		cluster.meanColor = colors[colorIndex].color
		clusters[i] = cluster
	}

	return clusters
}

func KMeans(colors []Color, k int) []Cluster {
	clusters := initClusters(colors, k)
	return clusters
}
