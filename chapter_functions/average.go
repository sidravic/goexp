package main

import "fmt"

func main(){

	numbersArray := [10]float64{2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 1.0, 0.0}
	fmt.Printf("\n numbers arrayis %v \n", numbersArray)
	result := Avg(&numbersArray)
	fmt.Printf("Average of %v is %v", numbersArray, result)
}

func Avg(numbers *[10]float64) (average float64){
	count := 0
	var sum float64

	for _, number := range *numbers{
		sum += number
		count += 1
	}

	average = sum/float64(count)
	return
}


