package main

import "fmt"
func main(){
	type Numbers struct{
		i int
		j int
	}

	numbers:=Numbers{1,2}

	fmt.Println(numbers,Numbers{3,4})
    
	num:=&numbers
	num.i=1
	fmt.Println(num)

}