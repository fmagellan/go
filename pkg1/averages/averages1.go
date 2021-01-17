package averages

// Average function
// Input: An array of integers

// Output: Average in a float64
func Average(numbers []int) float64 {
	var total int = 0
	for _, value := range numbers {
		total += value
	}

	return float64(total) / float64(len(numbers))
}
