package main

import (
	"fmt"
	"errors"
)

func max_mult(a, b, divider int) (int, error) {
	for i := b; i >= a; i-- {
		if i % divider == 0 {
			return i, nil
		}
	}
	return 0, errors.New("no such number")
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	answer, err := max_mult(a, b, 7)
	fmt.Print(answer, err)
}