package main

import (
	"errors"
	"fmt"
)

const divider = 7

func maxMult(a, b int) (int, error) {
	for i := b; i >= a; i-- {
		if i%divider == 0 {
			return i, nil
		}
	}
	return 0, errors.New("no such number")
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	answer, err := maxMult(a, b)
	fmt.Print(answer, err)
}
