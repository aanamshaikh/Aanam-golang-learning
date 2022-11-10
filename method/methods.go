package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Number struct{
	num1 int
	num2 int
}

func (number Number)nums() int{
return number.num1 * number.num2
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	n:= Number{1,2}
	fmt.Println(n.nums())
}
