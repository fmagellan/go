package main

import "fmt"
import "github.com/fmagellan/go/pkg1/averages"

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6}
    fmt.Println("Average of numbers:", averages.Average(numbers))
}
