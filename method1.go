package main

import "fmt"

type Student struct {
    id int
    name string
}

func (v Student) Print() {
    fmt.Printf("Id: %d, Name: %s.\n", v.id, v.name)
}

func main() {
    var venky Student
    venky.id = 10
    venky.name = "Venkat"
    fmt.Println(venky)
    venky.Print()
}
