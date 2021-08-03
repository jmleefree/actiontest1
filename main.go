package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func Multiply(x int, y int) int {
	return x * y
}

func Sub2(x int, y int) int {
	return x - y
}

func Sub3(x int, y int, z int) int {
	return x - y
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func customError() error {
	return MyError("an error occurred")
}

func main() {
	//test := 3
	fmt.Printf("result : %d\n", Sum(5, 6))
	fmt.Printf("result : %d\n", Sum(5, 8))
	fmt.Printf("result : %d\n", Sum(7, 8))
	fmt.Printf("result : %d\n", Sum(8, 8))
	fmt.Printf("result : %d\n", Sum(9, 8))

	customError() // UNCHECKED
}
