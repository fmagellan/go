package main

import "fmt"
import "time"

func pinger(channel chan<- string) {
    channel <- "one"
    channel <- "two"
    channel <- "three"
}

func ponger(channel chan<- string) {
    channel <- "ONE"
    channel <- "TWO"
    channel <- "THREE"
}

func printer(channel <-chan string) {
    for {
        fmt.Println(<-channel)
    }
}

func main() {
    var channel chan string = make(chan string)

    go pinger(channel)
    go ponger(channel)
    go printer(channel)

    time.Sleep(time.Second*15)
}
