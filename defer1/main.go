package main

import (
	"fmt"
)

func test1() (myInt int) {
    myInt = 1

    defer func() {
        myInt++
    }()

    return myInt
}

func test2() (myInt int) {
    myInt = 1

    defer func() {
        myInt++
    }()

    return func() int {
        fmt.Println(myInt)
        return myInt
    }() 
}

func main() {
    fmt.Printf("test1: %d\n", test1())
    fmt.Printf("test2: %d\n", test2())
}