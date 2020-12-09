package main

import "fmt"
import "time"

func main() {
    channel1 := make(chan string)
    channel2 := make(chan string)

    go func() {
        channel1 <- "one"
        channel1 <- "two"
        channel2 <- "THREE"
    }()

    go func() {
        channel2 <- "ONE"
        channel2 <- "TWO"
        channel2 <- "three"
    }()

    go func() {
        for {
            select {
            case msg := <-channel1:
                fmt.Println("Channel1:", msg)
            case msg := <-channel2:
                fmt.Println("Channel2:", msg)
            }
        }
    }()

    time.Sleep(time.Second * 10)
}
