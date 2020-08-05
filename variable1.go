// Copyright 2020 Magellan
package main

import "fmt"

func main() {
    var x int
    x = 10
    fmt.Println(x)
    fmt.Printf("x is of type %T\n", x)

    var y float32
    y = 11.0
    fmt.Println(y)
    fmt.Printf("y is of type %T\n", y)

    z := 12.0
    fmt.Println(z)
    fmt.Printf("z is of type %T\n", z);
}
