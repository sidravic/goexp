package main

import (
	"fmt"
	"math"
)

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string{
	return fmt.Sprintf("\n math: square root error \n", float64(f))
}

func (f *NegativeSqrtError) New(text) error{

}


func Sqrt(f float64) (float64, error) {
	if f < 0.0 {
		fmt.Println("Less than 0.0")
		return 0.0, NegativeSqrtError(f).Error()
	}
	return math.Sqrt(f), nil
	
}

func main() {
	val, err := Sqrt(-2)
	if(err != nil){
		fmt.Print(err)
	}
	
	fmt.Printf("%v \n", err)
	fmt.Printf("\n %v \n", val)
}