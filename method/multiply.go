package main

import "fmt"

type multiply int
type divide int

func (m multiply)multiplyTenTimes() multiply{

	return m*10
}

func (d divide)divideByTen() divide{

	return d/10
}

func main(){

var m multiply=10
fmt.Println(m.multiplyTenTimes())
 var d divide=10
 d.divideByTen()
}