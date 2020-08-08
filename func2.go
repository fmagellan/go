// Copyright 2020 Magellan
package main

import "fmt"

func main() {
    one, two := swap("ONE", "TWO")
    fmt.Println(one, two)
}

func swap(a, b string) (string, string) {
    return b, a
}
