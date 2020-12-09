package main

import "fmt"

type Dog struct {
    name string
}

type Cat struct {
    name string
}

func (dog Dog) says() string {
    return "Ruff"
}

func (cat Cat) says() string {
    return "Meow"
}

type Animal interface {
    says() string
}

func says(animal Animal) {
    fmt.Println(animal.says())
}

func main() {
    c := Cat{"Cat"}
    d := Dog{"Dog"}

    var animal Animal
    animal = c
    says(animal)
    animal = d
    says(animal)
}
