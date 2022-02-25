//случайными буквами заполнить матрицу.
// подсчитать количество повторяющихся символов

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := "a"
	b := "x"
	matrix := [][]string{}
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 10; j++ {
		slice := []string{}
		min := int(a[0])
		max := int(b[0])
		for i := 0; i < 10; i++ {
			number := min + rand.Intn(int((max-min))+1)
			slice = append(slice, string(number))
		}
		matrix = append(matrix, slice)
	}
	m := map[string]int{}
	for _, rows := range matrix {
		for _, letter := range rows {

			if _, ok := m[letter]; ok {
				m[letter]++

			} else {
				m[letter] = 1
			}

		}

	}
	for _, UniqueLettersCount := range matrix {

		fmt.Println(UniqueLettersCount)
	}

	fmt.Println(m)
}
