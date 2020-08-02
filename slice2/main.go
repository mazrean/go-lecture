package main

import (
	"fmt"
	"time"
)

func main() {
	num := 10000000
	times := 5

	start := time.Now()
	for i := 0; i < num; i++ {
		s := []int{}
		for j := 0; j < times; j++ {
			s = append(s, j)
		}
	}
	end := time.Now()
	fmt.Println("no cap:", end.Sub(start).Milliseconds(), "ms")

	start = time.Now()
	for i := 0; i < num; i++ {
		s := make([]int, 0, times)
		for j := 0; j < times; j++ {
			s = append(s, j)
		}
	}
	end = time.Now()
	fmt.Println("with cap:", end.Sub(start).Milliseconds(), "ms")
}
