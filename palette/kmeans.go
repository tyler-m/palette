package palette

import "math"

func distance(p Color, q Color) float64 {
	qDifference := q.color.R - p.color.R
	bDifference := q.color.B - p.color.B
	gDifference := q.color.G - p.color.G
	sumOfSquares := qDifference*qDifference + bDifference*bDifference + gDifference*gDifference
	return math.Sqrt(float64(sumOfSquares))
}

func KMeans(colors []Color, k int) {

}
