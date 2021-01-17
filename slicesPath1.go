package main

import "fmt"
import "bytes"

type pathType []byte

func main() {
    path := pathType("/users/home/fmagellan")
    fmt.Println(path)
    fmt.Printf("%s\n", path)
    fmt.Println("Length:", len(path))
    path.truncate()
    fmt.Printf("After truncating(value receiver. so not supposed to work.): %s\n", path)
    path.toUpper()
    fmt.Printf("%s\n", path)
}

func (pPath pathType) truncate() {
    lastDelimiterPos := bytes.LastIndex(pPath, []byte("/"))
    if (lastDelimiterPos >= 0) {
        pPath = pPath[0:lastDelimiterPos]
    }
}

func (path pathType) toUpper() {
    for index, value := range path {
        if 'a' <= value && value <= 'z' {
            path[index] = value + 'A' - 'a'
        }
    }
}
