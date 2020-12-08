package main

import "fmt"

func swap(pX, pY *int) {
    *pX, *pY = *pY, *pX
}

func main() {
    x, y := 1, 2
    fmt.Println(x, y)
    swap(&x, &y)
    fmt.Println(x, y)
}
