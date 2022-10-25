package main

//подсчет суммы чисел в слаисе
import "fmt"

func main() {
	s := []int{2, 8, 6, 9, 1, 7, 2}
	number := make(chan int)
	result := make(chan int)

	go rangeNumbers(s, number)
	go sumNumber(number, result)
	summ := <-result
	fmt.Println(summ)
}
func rangeNumbers(s []int, ch chan int) {
	for _, val := range s {
		ch <- val
	}
	close(ch)
}

func sumNumber(m, result chan int) {
	r := 0
	for val := range m {
		r += val
	}
	result <- r
}
