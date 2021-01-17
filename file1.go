package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("channel1.go");
    if err != nil {
        fmt.Println("File doesn't exist.")
        return
    }

    fmt.Println(string(data))

    toWrite := []byte("line three\nline four")
    toWrite = append(data, toWrite...)
    err = ioutil.WriteFile("temp.txt", toWrite, 0644)
    if err != nil {
        fmt.Println("File doesn't exist.")
        return
    }
}
