// Copyright 2020 Magellan
package main

import "fmt"

func main() {
    var i int
    var j = 20

    fmt.Printf("Loop 1: ")
    for i = 1; i <= 10; i++ {
        fmt.Printf("%d ", i)
    }

    fmt.Printf("\nLoop 2: ")
    for i <= j {
        fmt.Printf("%d ", i)
        i++
    }
    fmt.Printf("\n")

    numbers := [10] int { 21, 22, 23, 24, 25 }
    fmt.Printf("\nLoop 3: ")
    for index, element := range numbers {
        fmt.Printf("[%d]: %d ", index, element)
    }
    fmt.Printf("\n")
}
