package main

import "fmt"

func main(){
	var a int = 72
	var b int = 44

	
	var x, y int  = Sort(a, b)
	fmt.Printf("returned values %v and %v", x, y)
}

func Sort(x int, y int) (int, int){
	if x < y {
		return x, y
	}

	return y, x
}