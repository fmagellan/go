package main

import "fmt"

func main() {
    s1 := [] int {1, 2, 3, 4, 5}
    for index, element := range s1 {
        fmt.Println(index, element)
    }

    var index int
    var element string
    s2 := [] string {
        "one",
        "two",
        "three",
    }
    for index, element = range s2 {
        fmt.Println(index, element)
    }

    for index = range s2 {
        fmt.Printf("Index: %v\n", index)
    }

    for _, element = range s2 {
        fmt.Printf("Value: %v\n", element)
    }
}
