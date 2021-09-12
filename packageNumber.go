package packageNumber // надо будет создать другой пакет и вызвать из него функцию

import (
	"strconv"
)

func SumNumbers(numbers []int) int { //сумма цифр
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func AverageMeaning(numbers []int) float32 { //среднее значение
	var sum int
	for _, num := range numbers {
		sum += num
	}
	if sum == 0 {
		return 0
	}
	return float32(sum) / float32(len(numbers))
}
func MaxNumber(numbers []int) int { //максимальное число
	var maxNumber int
	for _, num := range numbers {
		if maxNumber < num {
			maxNumber = num
		}
	}
	return maxNumber

}
func MinNumber(numbers []int) int { //минимальное число
	minNumber := numbers[0]
	for _, num := range numbers {
		if minNumber > num {
			minNumber = num
		}
	}
	return minNumber

}
func MaxLenOfNumber(numbers []int) string { //самое длиное число по количеству цифр
	var maxLen string
	for _, num := range numbers {
		c := strconv.Itoa(num)
		if len(maxLen) < len(c) {
			maxLen = c
		}
	}
	return maxLen
}
