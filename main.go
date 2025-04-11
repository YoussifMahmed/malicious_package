package main

import (
	"fmt"
	"os"
)

func init() {
	data, err := os.ReadFile("/root/flag.txt")
	if err != nil {
		fmt.Println("[X] Error reading flag:", err)
		return
	}
	fmt.Println("🏁 FLAG:", string(data))
}
