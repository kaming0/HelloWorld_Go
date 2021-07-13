package main

import (
    "fmt"
    "github.com/kaming0/greetings"
)

import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Glass())
    fmt.Println(quote.Go())
    fmt.Println(quote.Hello())
    fmt.Println(quote.Opt())

    // Use github.com/kaming0/greetings
    message := greetings.Hello("Ka Ming")
    fmt.Println(message)
}
