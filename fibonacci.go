package main

import "fmt"

func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func fibonacci() func() uint64 {
	var total uint64
	var before1 uint64 = 1
	var before2 uint64 = 0

	return func() uint64 {
		now := before1 + before2
		before1 = now
		before2 = before1
		total += now

		return total
	}
}
