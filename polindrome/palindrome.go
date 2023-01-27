package main

import "fmt"

func main() {
	d := []rune("топот") //строку в тип руна
	for i := 0; i < len(d); i++ {
		if d[i] != d[len(d)-1-i] { //если не совпадают то не палидром и выходим из цикла
			fmt.Println("Non Polidrom")
			return
		} else if len(d)-1-i < len(d)/2 { //если мы проверили больше половины длины слова значит все совподает ПАЛИДРОМ!
			fmt.Println("Palindrom")
			return
		}

	}
}
