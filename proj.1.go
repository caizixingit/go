package main

import "fmt"

func main() {
	num := 1000

	// muti3 := (num - 1) / 3
	// muti5 := (num - 1) / 5

	// count3 := 0
	// count5 := 0
	// for i := 1; i <= muti3; i++ {
	// 	count3 += i
	// }

	// for i := 1; i <= muti5; i++ {
	// 	count5 += i
	// }
	total := 0
	for i := 0; i < num; i++ {
		if i%3 == 0 {
			total += i
			continue
		}
		if i%5 == 0 {
			total += i
			continue
		}
	}

	fmt.Println(total)
}
