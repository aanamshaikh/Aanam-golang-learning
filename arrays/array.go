package main

import "fmt"

func main(){
	var s =[2]int{} // array
    var a [10]string
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var b []int = primes[1:4]
	fmt.Println(b)
	
	fmt.Println(s,a)
}