package main

import (
	std "fmt"
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {
	std.Println(^1)

	std.Println(1 << 10)

	std.Println(1 >> 1)

	var c float64 = 1 >> 2
	std.Println(c)

	std.Println(KB)
	std.Println(MB)
	std.Println(GB)
	std.Println(TB)
}
