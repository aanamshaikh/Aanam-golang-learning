package main

import "fmt"

func main(){
	q:=[]int{23,4,5} 
	fmt.Println(q)

	r:=[]string{"Aanam","Tom","Harry"}
	fmt.Println(r)


	s:=[]struct{
		i int
		j bool
	}{
		{1,true},
	}

	fmt.Println(s)
}