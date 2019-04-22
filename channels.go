package main

import "fmt"

func main() {

    // Create a new channel with `make(chan val-type)`.
    // Channels are typed by the values they convey.
    messages := make(chan string)

    // _Send_ a value into a channel using the `channel <-`
    // syntax.
    go func() { messages <- "ping" }()

    // The `<-channel` syntax _receives_ a value from the
    // channel.
    fmt.Println(msg)
}

