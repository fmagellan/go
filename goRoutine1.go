package main

import "fmt"

func printNumbers(name string, counter int) {
    for index := 0 ;index < counter; index++ {
        fmt.Println(name, ":", index)
    }
}

func main() {
    var str string = "Five"
    go printNumbers(str, 5)
    str = "Ten"
    go printNumbers(str, 10)

//    fmt.Scanf("%d", &str)
}
