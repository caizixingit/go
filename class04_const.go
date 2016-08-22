package main

import (
	"fmt"
)

const mm = '3'
const (
	a = 1
	b = a + 2
	c = mm + 3
	d
	e = yy - 1
	f = iota
	g
)

const xx = iota

const yy = 4

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(xx)

}
