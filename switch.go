package main

import "fmt"

func main(){
	fmt.Printf(" status %v \n", shouldEscape('x'))
}

func shouldEscape(c byte) bool {
	switch c {
	case 'a', 'b', 'c', 'd':
		return true
	}
	return false
}