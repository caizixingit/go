package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	a = 9
	doA(&a)
	fmt.Println(a)
	doB(a)
	fmt.Println(a)
}

func doA(a *int) {
	b := a
	*b = 5
	fmt.Println(*b)
}

func doB(a int) {
	b := &a
	*b = 8
	fmt.Println(*b)
}
