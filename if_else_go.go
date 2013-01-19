package main

import "fmt"

func main(){
	if true {
		fmt.Printf("This is here")
	} else {
		fmt.Printf("Here")
	}


	for i:= 0 ; i < 10; i++ {
		if(i == 5){
			break;
		}
		fmt.Printf("This is the %v\n", i)
		println("This is the second text \n", i)
	}
}