package main

import "fmt"

func main(){
	numbers := []float64{1.3, 2.0, 3.0, 4.0, 7.0}	
	fmt.Printf("%v \n", numbers)
	result := Avg(numbers[:])
	fmt.Printf("result %v \n", result)
}

func Avg(numbers []float64) (avg float64){
	var sum float64 = 0.0
	var count int = 0

	for _, number := range numbers{
		fmt.Printf("- %v \n ", number )
		sum += number
		count += 1
	}
	avg = sum/(float64(count))
	return
}	