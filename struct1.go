package main

import "fmt"

func main() {
    emp := struct {
        firstName string
        lastName string
        age int
    }{
        firstName: "Ferdinand",
        age: 45,
        lastName: "Magellan",
    }

    fmt.Println(emp)
    fmt.Printf("Type of the struct: %T.\n", emp)
}
