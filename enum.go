package main

// When you run the command, go generate enum.go,
// it will run the below go-generate command from enum.go file.
// And will generate a file called number_string.go.
// number_string.go will define a receiver method, that converts Enum-value to a string,
// as part of the current package.
// Now execute the program as "go run enum.go number_string.go"

import "fmt"

//go:generate stringer -type=Number
type Number int

const (
    One Number = iota
    Two
    Three
    Four
    Five
)

func main() {
    fmt.Println(Four)
}
