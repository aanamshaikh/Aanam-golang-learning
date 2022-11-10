package main

import "fmt"

type Number int

func (n Number) num() int {
	return int (n*n)
}
func main() {

	fmt.Println(Number.num(25))
}
