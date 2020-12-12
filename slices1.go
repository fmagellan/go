package main

import "fmt"

func main() {
    var arr5 [5]int = [5]int{5, 4, 3, 2, 1}
    fmt.Println(arr5)

    var arr3 [3]int = arr5[0 : 3]
    fmt.Println(arr3)

    var arr4 []int = arr5[0 : 4]
    fmt.Println(arr4)
}
