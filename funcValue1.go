// Copyright 2020 Magellan
package main

import "fmt"

func main() {
    squareFunc := func(x int) int {
        return (x * x)
    }

    fmt.Println("Square of 10:", squareFunc(10))
    fmt.Println("Square of 10:", squareFunc(11))
}
