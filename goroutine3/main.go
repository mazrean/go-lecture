package main

import "fmt"

func main() {
	times := 5

	ch := make(chan int, times-2) //容量times-2のchannel作成
	for i := 0; i < times; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}

	for i := 0; i < times; i++ {
		fmt.Printf("len:%d\n", len(ch))
		val := <-ch
		fmt.Printf("len: %d, val: %d\n", len(ch), val)
	}

	fmt.Print("end")
}
