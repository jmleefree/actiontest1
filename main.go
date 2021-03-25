package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("result : %d\n", Sum(5, 6))
}
