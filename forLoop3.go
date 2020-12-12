package main

import "fmt"

func main() {
    var arr [5]string = [5]string{
        "zero",
        "one",
        "two",
        "three",
        "four",
    }

    fmt.Println("=============")
    for index, value := range arr {
        fmt.Printf("arr[%v]: %v\n", index, value)
    }

    fmt.Println("=============")
    for _, value := range arr {
        fmt.Println(value)
    }

    fmt.Println("=============")
    for index, _ := range arr {
        fmt.Println(index)
    }
}
