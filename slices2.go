package main

import "fmt"

func main() {
    arr5 := [5] string {
        "one",
        "two",
        "three",
        "four",
        "five",
    }

    fmt.Println(arr5)

    arr02 := arr5[0 : 3]
    fmt.Println(arr02)

    arr24 := arr5[2 : 5]
    fmt.Println(arr24)

    arr02[2] = "arr02"
    fmt.Println(arr5, arr02, arr24)

    arr24[0] = "arr24"
    fmt.Println(arr5, arr02, arr24)

    return
}
