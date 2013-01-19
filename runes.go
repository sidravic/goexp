package main
import "fmt"

func main(){
	x := "Siddharth"
	c := []rune(x)
	c[0] = "Ps"
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}