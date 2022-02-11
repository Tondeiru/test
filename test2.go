package main

import (
	"fmt"
	"time"
)

func sumSlices(x []float64, y []float64) float64 {

	totalx := 0.0
	for _, valuex := range x {
		go func() {
			totalx += valuex
		}()
		time.Sleep(5 * time.Millisecond)
	}

	totaly := 0.0
	for _, valuey := range y {
		go func() {
			totaly += valuey
		}()
		time.Sleep(5 * time.Millisecond)
	}

	return totalx + totaly
}

func main() {
	x := []float64{1.3, 2, 3, 4, 5}
	y := []float64{11, 12, 130, 14, 15}

	fmt.Println(sumSlices(x, y))
}
