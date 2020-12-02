package main

import "fmt"

func main() {
    var s []int = make([]int, 0, 5)
    fmt.Printf("Len:%d, Capacity:%d, s:%v\n", len(s), cap(s), s)

    s = append(s, 0)
    fmt.Printf("Len:%d, Capacity:%d, s:%v\n", len(s), cap(s), s)

    s = append(s, 1)
    fmt.Printf("Len:%d, Capacity:%d, s:%v\n", len(s), cap(s), s)
}
