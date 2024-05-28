package calculation

func Average(numbers []float64) float64 {
	lenArr := len(numbers)
	total := 0.0
	for _, val := range numbers {
		total += val
	}
	mean := total / float64(lenArr)
	return mean
}