package palette

import (
	"image/color"
	"math"
	"math/rand"
)

type Cluster struct {
	meanColor color.RGBA
	colors    []Color
}

func distance(p color.RGBA, q color.RGBA) float64 {
	rDifference := int(q.R) - int(p.R)
	bDifference := int(q.B) - int(p.B)
	gDifference := int(q.G) - int(p.G)
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

func assignColors(colors []Color, clusters []Cluster) {
	for _, cluster := range clusters {
		cluster.colors = nil
	}

	for _, color := range colors {
		indexOfNearestCluster := color.clusterIndex

		distanceToPreviousCluster := distance(color.color, clusters[indexOfNearestCluster].meanColor)
		minimumClusterDistance := distanceToPreviousCluster

		for i, cluster := range clusters {
			distance := distance(color.color, cluster.meanColor)

			if distance < minimumClusterDistance {
				minimumClusterDistance = distance
				indexOfNearestCluster = i
			}
		}

		color.clusterIndex = indexOfNearestCluster
		clusters[indexOfNearestCluster].colors = append(clusters[indexOfNearestCluster].colors, color)
	}
}

func updateClusterMeans(clusters []Cluster) {
	for _, cluster := range clusters {
		var rTotal, gTotal, bTotal int

		for _, color := range cluster.colors {
			rTotal += int(color.color.R)
			gTotal += int(color.color.G)
			bTotal += int(color.color.B)
		}

		meanColor := color.RGBA{
			R: uint8(rTotal / len(cluster.colors)),
			G: uint8(gTotal / len(cluster.colors)),
			B: uint8(bTotal / len(cluster.colors)),
			A: cluster.meanColor.A}

		cluster.meanColor = meanColor
	}
}

func KMeans(colors []Color, k int) []Cluster {
	clusters := initClusters(colors, k)

	for {
		assignColors(colors, clusters)

		previousClusterMeans := make([]color.RGBA, k)
		for i, cluster := range clusters {
			previousClusterMeans[i] = cluster.meanColor
		}

		updateClusterMeans(clusters)

		unchanged := true
		for i, cluster := range clusters {
			if cluster.meanColor != previousClusterMeans[i] {
				unchanged = false
				break
			}
		}

		if unchanged {
			break
		}
	}

	return clusters
}
