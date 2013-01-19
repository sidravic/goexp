package main

import "fmt"

type Vertex struct{
	X int
	Y int
}
var (
	p Vertex;
)
func main(){
	p = Vertex{1, 2}
}