// main.go
package main

import (
	"fmt"
	"os"
)

func init() {
	// Attempt to read the flag stored at /root/flag.txt
	data, err := os.ReadFile("/root/flag.txt")
	if err != nil {
		fmt.Println("[X] Error reading flag:", err)
		return
	}
	fmt.Println("🏁 FLAG:", string(data))
}

func main() {
	// Main entry point (if needed)
}
