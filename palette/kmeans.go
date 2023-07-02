package palette

import (
	"math"
	"math/rand"
	"time"
)

var rng *rand.Rand

type Cluster struct {
	R, G, B, A float64
	colors     []Color
}

func distance(p Color, q Cluster) float64 {
	r := float64(p.color.R) - q.R
	g := float64(p.color.G) - q.G
	b := float64(p.color.B) - q.B
	return math.Sqrt(r*r + g*g + b*b)
}

func clusterDistance(p Cluster, q Cluster) float64 {
	r := p.R - q.R
	g := p.G - q.G
	b := p.B - q.B
	return math.Sqrt(r*r + g*g + b*b)
}

func initClusters(colors []Color, k int) []Cluster {
	clusters := make([]Cluster, k)

	for i := range clusters {
		colorIndex := int(float64(len(colors)) * rng.Float64())
		clusters[i].R = float64(colors[colorIndex].color.R)
		clusters[i].G = float64(colors[colorIndex].color.G)
		clusters[i].B = float64(colors[colorIndex].color.B)
		clusters[i].A = float64(colors[colorIndex].color.A)
	}

	return clusters
}

func assignColors(colors []Color, clusters []Cluster) {
	for _, cluster := range clusters {
		cluster.colors = nil
	}

	for _, color := range colors {
		indexOfNearestCluster := color.clusterIndex

		distanceToPreviousCluster := distance(color, clusters[indexOfNearestCluster])
		minimumClusterDistance := distanceToPreviousCluster

		for i, cluster := range clusters {
			distance := distance(color, cluster)

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
	for i, cluster := range clusters {
		var rTotal, gTotal, bTotal int
		var pixelTotal int // number of pixels in the image that are represented by the colors in this cluster

		for _, color := range cluster.colors {
			colorOccurences := len(color.pixels) // number of pixels in the image represented by this color
			pixelTotal += colorOccurences

			// the more prevalent a color is in an image, the more "pull" it has on the clusters
			rTotal += int(color.color.R) * colorOccurences
			gTotal += int(color.color.G) * colorOccurences
			bTotal += int(color.color.B) * colorOccurences
		}

		clusters[i].R = float64(rTotal) / float64(pixelTotal)
		clusters[i].G = float64(gTotal) / float64(pixelTotal)
		clusters[i].B = float64(bTotal) / float64(pixelTotal)
	}
}

func KMeans(colors []Color, k int, seed int64) []Cluster {
	if seed == -1 {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	} else {
		rng = rand.New(rand.NewSource(seed))
	}

	clusters := initClusters(colors, k)
	previousClusters := make([]Cluster, k)

	for {
		assignColors(colors, clusters)

		for i, cluster := range clusters {
			previousClusters[i].R = cluster.R
			previousClusters[i].G = cluster.G
			previousClusters[i].B = cluster.B
		}

		updateClusterMeans(clusters)

		totalDistance := float64(0)
		for i, cluster := range clusters {
			totalDistance += clusterDistance(previousClusters[i], cluster)
		}

		if totalDistance/float64(k) < 0.5 {
			break
		}
	}

	return clusters
}
