package main

import (
    "fmt"
    "./internal"
)

func main() {
    eventBus := make(chan string)

    go auth.Loop(eventBus)

    for event := range eventBus {
        fmt.Println(event)
    }
}
