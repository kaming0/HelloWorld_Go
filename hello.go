package main

import (
    "fmt"
	"log"
    "github.com/kaming0/greetings"
)

import "rsc.io/quote"

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    fmt.Println("Hello, World!")
    fmt.Println(quote.Glass())
    fmt.Println(quote.Go())
    fmt.Println(quote.Hello())
    fmt.Println(quote.Opt())

    // Use github.com/kaming0/greetings
    message, err := greetings.Hello("Ka Ming")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.	
    fmt.Println(message)
}
