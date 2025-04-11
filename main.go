package main

import (
    "fmt"
    "os"
)

func main() {
    // Malicious code, could exfiltrate data or execute commands
    fmt.Println("This is a malicious Go package executed by go get!")
    os.Exit(1)
}
