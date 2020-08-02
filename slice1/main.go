package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4}
	s := arr[0:3]
	s = append(s, 10)
	fmt.Println(arr, s)
}
