package main

import "fmt"

func myFunc(b int, words []string) (status bool, data []string){
	fmt.Printf("b message received %v", b)

	for index, word := range words{
		fmt.Printf("%d) %v ", index, word)
		fmt.Printf("\n")
	}

	 status = true
	 data[0] = "shit" 
	return
}

func main(){
	words := make([]string, 10)
	words[0] = "Siddharth"
	words[1] = "Priyanka"
	words[2] = "Piper"
	words[3] = "At"
	words[4] = "gates"
	words[5] = "of"
	words[6] = "dawn"
	status, stuff := myFunc(22, words)
	if(status){
		fmt.Printf("================%v", stuff)
	}
}