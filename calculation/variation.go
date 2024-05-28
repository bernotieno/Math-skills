package calculation

// Calculates the variance of the slice by also calling the average function
func Variance(numbers []float64) float64 {
	lenArr := len(numbers)
	var sum float64
	for _, val := range numbers {
		sum += (val - Average(numbers)) * (val - Average(numbers))
	}
	return sum / float64(lenArr)
}
