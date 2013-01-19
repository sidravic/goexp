package main

import "fmt"
import "strconv"

func main(){
	var a [3]int;

	a[0] = 22
	a[2] = 33
	a[1] = 0

	b := [3]int{1,2}
	c := [2][2]string{[2]string{"Siddharth", "Piyoo"}, [2]string{"28","28"}}

	fmt.Printf("%v \n", a)
	fmt.Printf("Ths second one %v \n", b)
	fmt.Printf("This multi-dimensional %v \n", c)
	i, _ := strconv.ParseInt(c[1][1], 10, 64)
	

	fmt.Printf("Val %v \n", i)
}