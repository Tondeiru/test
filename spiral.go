package main

import (
	"fmt"
)

const N = 5

func spiralka(n int) []int {
	left, top, right, bottom := 0, 0, N-1, N-1
	size := N * N
	s := make([]int, size)
	step := 1
	for left < right {
		//napravo
		for r := left; r <= right; r++ {
			s[top*N+r] = step
			step++
		}
		top++
		// sprava vniz
		for b := top; b <= bottom; b++ {
			s[b*N+right] = step
			step++
		}
		right--
		if top == bottom {
			break
		}
		// vlevo
		for l := right; l >= left; l-- {
			s[bottom*N+l] = step
			step++
		}
		bottom--
		// vverh
		for t := bottom; t >= top; t-- {
			s[t*N+left] = step
			step++
		}
		left++
	}
	//centr

	return s
}

func main() {
	width := 2 // shirina mejdu stolbcami
	for position, draw := range spiralka(N) {
		fmt.Printf("%*d ", width, draw)
		if position%N == N-1 {
			fmt.Println("")
		}
	}
}
