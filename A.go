package main

import "fmt"
import "unicode/utf8"

func main(){
	for i:=1;i<=100;i++{
		for j:=1; j<=i; j++ {
			fmt.Printf("A")	
		}
		fmt.Printf("\n")
	}

	// Characters in a String

	var placeholder string = "asSASA ddd dsjkdsjs dk"
	fmt.Printf("\n")	
	fmt.Printf("Lenght of the string %v \n",len(placeholder))

	byteCount := utf8.RuneCount(placeholder)
	fmt.Printf("Byte Count is %v \n", byteCount)

}

