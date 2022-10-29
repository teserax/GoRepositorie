// Дан массив целых чисел и ещё одно целое число.
// Удалите все вхождения этого числа из массива
// (пропусков быть не должно).
// РЕАЛИЗОВАТЬ  ЧEРЕЗ ФУНКЦИЮ
package main

import (
	"fmt"
)

func removalNumber(n int, s []int) []int {
	sl := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] != n {
			sl = append(sl, s[i])
		}

	}
	return sl

}

func main() {
	slice := []int{8, 2, 2, 2, 6, 2, 9, 4, 2, 7, 3, 2}
	removal := 2

	slice = removalNumber(removal, slice)
	fmt.Println(slice)

}
