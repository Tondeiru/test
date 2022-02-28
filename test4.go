package main

import (
	//"fmt"
	"math/rand"
	"time"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99) - rand.Intn(99)

	}
	return slice
}

func selection_sort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		// Find the minimum element in unsorted array
		min_idx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}

		// Swap the found minimum element with the first element
		tmp := arr[i]
		arr[i] = arr[min_idx]
		arr[min_idx] = tmp
	}
}

func bubble_sort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {

			if arr[j-1] > arr[j] {
				median := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = median
			}
		}

	}
}

func merge_sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := merge_sort(arr[:len(arr)/2])
	second := merge_sort(arr[len(arr)/2:])
	return merge(first, second)

}

func merge(a []int, b []int) []int {

	final := []int{}
	j := 0
	i := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func quick_sort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	//random pivot
	pivot := rand.Int() % len(arr)

	//pivot to the right
	arr[pivot], arr[right] = arr[right], arr[pivot]

	//elements smaller than pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	//pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	quick_sort(arr[:left])
	quick_sort(arr[left+1:])

}

func main() {

	tmp := generateRandomSlice(100000)
	//fmt.Println(tmp)

	//selection_sort(tmp)
	//fmt.Println(tmp)

	//bubble_sort(tmp)
	//fmt.Println(tmp)

	//merge_sort(tmp)
	//fmt.Println(tmp)

	//quick_sort(tmp)
	//fmt.Println(tmp)
}
