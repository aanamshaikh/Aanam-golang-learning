package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[int]Vertex
var n map[string]Vertex

func main() {
	m = make(map[int]Vertex)
	m[1] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m[1])

	n = make(map[string]Vertex)
	n["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(n["Bell Labs"])
}
