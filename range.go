package main

import "fmt"

func main(){	

	//list := []string{"Welcome to the winter of my dicontent", "TEst"}

	for k, v := range ["a", "b"] {
		fmt.Printf("Some Shit %v %c \n", k, v)
	}
}	