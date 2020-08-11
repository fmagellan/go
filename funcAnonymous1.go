// Copyright 2020 Magellan
package main

import "fmt"

func getSequence(start int) func(int) int {
    return func(increment int) int {
        start+=increment
        return start
    }
}

func main() {
    func1 := getSequence(10)
    fmt.Println(func1(1))
    fmt.Println(func1(2))

    func2 := getSequence(20)
    fmt.Println(func2(2))
    fmt.Println(func2(3))
}
