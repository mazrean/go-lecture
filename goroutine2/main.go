package main

import "fmt"

func main() {
	times := 5

	ch := make(chan int, times-2) //容量5のchannel作成
	for i := 0; i < times; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}

	for len(ch) < times { //chに値がたまりきるまで待つ
		fmt.Println(len(ch))
	}

	fmt.Print("end")
}
