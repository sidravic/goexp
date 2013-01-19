package main

import "fmt"


func main(){
	x := "Hello World"
	xx := []rune(x)

	byteText := []byte(x)
	fmt.Printf("Some Data %v \n", xx)



	for _, v := range(xx) {
		fmt.Printf("%c \n", v)
	}

	fmt.Printf("%v \n", byteText)
	// t := 'test'

	// fmt.Printf('%v \n', t)

}