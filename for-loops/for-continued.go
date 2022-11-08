package main

import "fmt"

func main() {
	sum := 1
	for  sum < 1000 { // or  for ; sum < 1000;
		sum += sum
	}
	fmt.Println(sum)

	for {
		println("This is an infinite loop")
	}
}
