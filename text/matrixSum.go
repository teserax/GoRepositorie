//Заполнить матрицу случайными числами.
//На главной диагонали разместить суммы
//элементов, которые лежат на той же
//строке и том же столбце.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const N = 10
	const R = 10
	m := [N][R]int{}
	rand.Seed(time.Now().UnixNano())
	sumHorizont := make(map[int]int) //мапа для суммы горизантальной строки
	for i := 0; i < N; i++ {

		for j := 0; j < R; j++ {
			num := rand.Intn(20) + 1 //случайное число
			m[i][j] = num
			sumHorizont[i] += num //записываем сумму чисел в горизантальной строке

		}
	}
	sumVert := make(map[int]int) //мапа для суммы вертикальной строки

	for _, row := range m {
		for sum, num := range row {
			sumVert[sum] += num //записываем сумму чисел в вертикальной строки
		}
		for i := 0; i < len(m); i++ {
			m[i][i] = sumHorizont[i] + sumVert[i] //добавляем сумму строк в матрицу итерируясь по диагонали
		}

	}
	fmt.Println("")

	for _, row := range m {
		fmt.Printf("%3d\n", row)
	}

}
