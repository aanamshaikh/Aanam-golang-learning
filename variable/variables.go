package main

import "fmt"

var c, python, java bool //by default false

// variables with initializers
var a, b int = 1, 2 


func main() {

	var i int //by default false	
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!" // variables with initializers
	fmt.Println(a, b, c, python, java)

	//short variable declaration
	var k, l int = 1, 2
	m := 3
	// z, go, ruby := true, false, "no!"
	fmt.Println(k,l,m)
}


