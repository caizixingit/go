/**
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
package main

import "fmt"

func main() {
	list := make(map[int]int)
	var tmp int
	total := 1
	prefix := 0

	for i := 1; i < 21; i++ {
		tmp = i
		fmt.Println(list)
		for _, j := range list {
			if tmp%j == 0 {
				tmp = tmp / j
			}
		}

		if tmp > 1 {
			total *= tmp
			list[prefix] = tmp
			prefix++
		}
	}

	fmt.Println(total)
}
