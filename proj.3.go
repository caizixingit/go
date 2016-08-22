/**
 * The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	max := 600851475143
	num := 0

	for i := 2; i < int(math.Sqrt(float64(max))); i++ {
		if max%i == 0 {
			if isPrime(i) {
				num = i
			}
		}
	}

	fmt.Println(num)
}

func isPrime(num int) bool {
	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
