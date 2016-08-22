package main

import (
	"fmt"
)

func main() {
	a := 1
	var p *int = &a

	fmt.Println(p)
	fmt.Println(*p)
}
