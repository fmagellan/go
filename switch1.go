package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i % 15 == 0 {
            fmt.Println(i, ": FizzBuzz")
        } else if i % 5 == 0 {
            fmt.Println(i, ": Buzz")
        } else if i % 3 == 0 {
            fmt.Println(i, ": Fizz")
        } else {
            fmt.Println(i)
        }
    }

    for i := 1; i < 20; i++ {
        switch i % 3 {
        case 0:
            fmt.Println(i, ": Remainder 0")
        case 1:
            fmt.Println(i, ": Remainder 1")
        case 2:
            fmt.Println(i, ": Remainder 2")
        }
    }
}
