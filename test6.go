package main

import (
	"fmt"
)

func printFibonacci(n int) int {

	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, b+a
	}

	return a

}

func main() {

	var n = 3

	for i := 0; true; i++ {
		if n == printFibonacci(i) {
			fmt.Print(i)
			return
		}

	}
	fmt.Print(-1)
}
