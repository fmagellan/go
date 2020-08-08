// Copyright 2020 Magellan
package main

import "fmt"

func main() {
    a := 2
    b := 10
    c := max(a, b)
    fmt.Println("Max is", c)
}

func max(a, b int) int {
    var result int
    if (a > b) {
        result = a
    } else {
        result = b
    }

    return result
}
