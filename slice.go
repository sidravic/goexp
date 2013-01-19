package main

import "fmt"

func main(){
	array := [10]int{1,2,3,4,5,6,7,8,9,0}

	array2 := [5]int{0,1,2,3,4}
	sl := make([]int, 7)

	c1 := copy(sl, array2[0:])
	fmt.Printf("== Data %v ==\n ", c1)
	fmt.Printf("== Data %v ==\n ", sl)
	c2 := copy(sl, array[7:])	
	fmt.Printf("== Data Count %v == \n", c2)
	fmt.Printf("== Data %v == \n", sl)

	s1 := array[:]	

	directSlice := make([]int, 10)
	directSlice[0] = 22
	s2 := append(directSlice, 11)
	

	fmt.Printf("Direct Slice %v \n", (s2))

	fmt.Printf("array size is %v \n", len(s1))
}