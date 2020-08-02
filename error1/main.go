package main

import (
	"errors"
	"fmt"
)

func div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("zero div") //yが0のときにエラーを返す
	}

	return x / y, nil //エラーがでなければnilを返す
}

func main() {
	result, err := div(2, 1)
	if err != nil {
		panic(err) // errがnil出ない場合、ここでpanic
	}

	fmt.Println(result)

	result, err = div(2, 0)
	if err != nil {
		panic(err)
	}
}
