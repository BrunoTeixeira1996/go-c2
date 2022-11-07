package main

import (
    "fmt"
    "os"
)

// Function that handles the errors
func run() error {
    fmt.Println("Hello")
    return nil
}

func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
