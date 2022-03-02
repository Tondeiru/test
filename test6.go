package main

import (
	"fmt"
	//"math"
)

func fibr(n int) int {
	if n < 2 {
		return 1
	}
	return fibr(n-2) + fibr(n-1)
}

func main() {

	var n = 3
	//fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if n == fibr(i) {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
