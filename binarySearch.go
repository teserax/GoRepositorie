//binary searh

package main

import (
	"fmt"
)

func binarySearch(s []int, item int) {
	start := 0
	end := len(s)
	middle := 0
	found := false
	position := -1
	count := 0
	for found == false && start <= end {

		middle = (start + end) / 2

		if s[middle] == item {
			found = true
			position = middle

		}
		if item < s[middle] {
			end = middle - 1

		} else {
			start = middle + 1
		}
		count++
	}
	fmt.Println(position, count)

}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7, 7, 8, 9}
	binarySearch(slice, 6)

}
