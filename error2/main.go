package main

import (
	"errors"
	"fmt"
)

func div(x int, y int) (int,error) {
	if y == 0 {
		return 0, errors.New("zero div") //yが0のときにエラーを返す
	}

	return x/y, nil //エラーがでなければnilを返す
}

func repeatDiv(num int, x int, y int) (int, error) {
	var err error
	for i:=0;i<num;i++ {
		a := y
		y,err = div(x,y)
		if err != nil {
			return 0, fmt.Errorf("failed to dividing(i: %d, x: %d, y: %d): %w", i, x, a, err)
		}
		x = a
	}

	return x, nil
}

func main() {
	result, err := repeatDiv(4, 5, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}