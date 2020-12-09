package main

import "fmt"
import "averages"

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6}
    fmt.Println("Average of numbers:", averages.Average(numbers))
}
