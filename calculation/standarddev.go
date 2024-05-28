package calculation

import "math"

// calculate the standard deviation by finding the squareroot of the variance
func StandDev(variance float64) float64 {
	// Find the square root of variance
	return math.Sqrt(variance)
}