package main

import (
	"fmt"
)

func despair() {
	for {
		fmt.Println("Help me!")
	}
}

func main() {
	go despair()
}