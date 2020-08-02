package main

import (
	"errors"
	"fmt"
)

type ErrZeroDivide struct {
	x int
}

func (e ErrZeroDivide) Error() string {
	return fmt.Sprintf("zero dividing(x:%d)", e.x)
}

var zeroDivide ErrZeroDivide

func div(x int, y int) (int, error) {
	if y == 0 {
		zeroDivide = ErrZeroDivide{
			x: x,
		}
		return 0, zeroDivide //yが0のときにエラーを返す
	}

	return x / y, nil //エラーがでなければnilを返す
}

func repeatDiv(num int, x int, y int) (int, error) {
	var err error
	for i := 0; i < num; i++ {
		a := y
		y, err = div(x, y)
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
		fmt.Println(errors.Is(err, zeroDivide))

		z := &ErrZeroDivide{}
		if errors.As(err, z) {
			fmt.Println(z.x)
		}

		panic(err)
	}

	fmt.Println(result)
}
