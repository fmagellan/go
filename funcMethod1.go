// Copyright 2020 Magellan
package main

import "fmt"

type Rectangle struct {
    length, breadth int
}

func (r Rectangle) area() int {
    return (r.length * r.breadth)
}

func main() {
    r := Rectangle{length:10, breadth:5}
    fmt.Println("Area:", r.area())
}
