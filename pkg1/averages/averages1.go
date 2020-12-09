package averages

func Average(numbers []int) float64 {
    var total int = 0
    for _, value := range numbers {
        total += value
    }

    return float64(total) / float64(len(numbers))
}
