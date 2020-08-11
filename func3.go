// Copyright 2020 Magellan
package main

import "fmt"

func main() {
    a, b := swap("A", "B")
    fmt.Println(a, b)
}

func swap(a, b string) (string, string) {
    return b, a
}
