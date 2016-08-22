package main

import (
	"fmt"
)

func demo() {
	a := 2
	if a := 1; a > 0 {
		fmt.Println(a)
	}

	fmt.Println(a)
}

func demo2() {
	i := 1
	for {
		i++
		if i > 3 {
			break
		}

		fmt.Println(i)
	}
}

func demo3() {
	i := 1
	for i <= 3 {
		i++
		fmt.Println(i)
	}
}

func demo4() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}

func demo5() {
	a := 10
	switch a {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("others")
	}
}

func demo6() {
	a := 10
	switch {
	case a > 0:
		fmt.Println(a)
		fallthrough
	case a > 1:
		fmt.Println(a)
	case a > 2:
		fmt.Println(a)
	default:
		fmt.Println("others")
	}
}

func demo7() {
LABEL1:
	for {
		for i := 1; i < 10; i++ {
			fmt.Print(i)
			if i > 3 {
				break LABEL1 //同理continue
			}
		}
	}

	goto LABEL2
	fmt.Println("OK")
LABEL2:
	fmt.Println("jump OK")
}

func demo8() {
LABEL:
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		continue LABEL
	}
}

func main() {
	//demo()
	//demo2()
	// demo3()
	// demo4()
	// demo5()
	// demo6()
	// demo7()
	demo8()
}
