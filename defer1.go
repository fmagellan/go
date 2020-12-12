package main

import "fmt"

func main() {
    x := 10
    defer fmt.Println("First argument", x)
    x++
    defer fmt.Println("Second argument", x)
    x++
}
