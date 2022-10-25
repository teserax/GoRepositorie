//заполнить матрицу построчно от начала и до конца

package main

import "fmt"

func main() {
	const N = 5
	const R = 5
	m := [N][R]int{}
	var ver, hor, count int
	count = 1
	for count <= N*R {

		m[ver][hor] = count
		if ver == 0 && hor%2 == 0 {
			hor++
			count++
			m[ver][hor] = count
			for ver == 0 {
				hor--
				ver++
				m[ver][hor] = count
				count++
				m[ver][hor] = count
				ver++
				fmt.Println(count)
			}
		} else if hor == 0 && ver%2 == 0 {

			m[ver][hor] = count
			for hor == 0 {
				if ver != 0 {
					ver--
				}
				hor++
				count++
				m[ver][hor] = count
				hor++
			}
			fmt.Println(ver)
		}

		count++
	}

	for _, val := range m {
		fmt.Printf("%2d\n", val)
	}

}
