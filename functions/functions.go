package main

import "fmt"

func add(x int, y int) int { // or add(x, y int)
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap(1, 2)
	fmt.Println(a, b)
}
