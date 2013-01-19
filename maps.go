package main

import "fmt"

func main() {
	monthDays := map[string]int{
		"jan": 31, "feb":29, "mar":31, "apr":30,
	}

	fmt.Printf("Days in feb %v \n", monthDays["feb"])
	//array1 := [10]int{}
	v, ok := monthDays["jan"]

	fmt.Printf("%v \n", v)
	fmt.Printf("%v \n", ok)
	for month , days := range(monthDays){
		fmt.Printf("The number of days for %v is %v \n", month, days)
	}

	delete(monthDays, "apr")

	fmt.Printf("%v \n", monthDays)

}