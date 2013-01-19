package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v \n", i)
	}


// Loops with goto
fmt.Printf("Loops with goto...\n")
x := 0
var array [10]int
Callback:	
	fmt.Printf("%v \n", x)
	array[x] = x * 2

	x += 1
	if( x < 10){
		goto Callback
	}

fmt.Printf("array is %v \n", array)


// FizzBuzz

for z := 0; z <=100; z++{
	switch{
	case z % 5 == 0 && z % 3 == 0:
		fmt.Printf("FizzBuzz \n")
	case z % 3 == 0:
		fmt.Printf("Fizz \n")
	case z % 5 == 0:
		fmt.Printf("Buzz \n")
	default:
		fmt.Printf("%v \n", z)
	}
}
	
}




