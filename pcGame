// Используя пять вариантов наборов карт 1)«6», «7», «8» 2)«7», «8», «9» 3)
// «6», «9», «10» 4)«6», «9», «8» 5)«7», «6», «10», сыграйте с компьютером.
// Введите с клавиатуры свой вариант и сравните с вариантом компьютера,
// который создайте, используя функцию генератора случайных чисел
// (три цифры от 6 до 10 включительно без повторения цифры в варианте). Если
// сумма цифр вашего варианта больше суммы цифр компьютера, вы выиграли.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	human := [][]int{{6, 7, 8}, {7, 8, 9}, {6, 9, 10}, {6, 9, 8}, {7, 6, 10}}
	choice := 0
startGame:
	fmt.Println("Welcome to our game")
	for {
		fmt.Println("Please select a number from 1 to 5 inclusive")
		fmt.Fscan(os.Stdin, &choice)

		if choice <= 0 || choice > 5 {
			fmt.Println("your number is incorrect please try again")
		} else {
			break
		}

	}
	humanChoice := 0

	for _, num := range human[choice-1] {

		humanChoice += num
	}
	pc := []int{}
	pcVal := 0
	m := map[int]int{}
	for i := 0; i < 3; i++ {
		num := rand.Intn(5) + 5
		if _, ok := m[num]; !ok {
			m[num]++
			pc = append(pc, num)
			pcVal += num
		}

		if len(m) < 3 && i >= 2 {
			i--
		}

	}
	fmt.Println("--------------------------------------")
	fmt.Printf("The amount of your choice is | %5d\nThe amount of PC choice is   | %5d\n", humanChoice, pcVal)
	if humanChoice > pcVal {
		fmt.Println("--------------------------------------")
		fmt.Println("Our congratulations you won")
		fmt.Println("--------------------------------------")
	} else {
		fmt.Println("--------------------------------------")
		fmt.Println("Unfortunately you lost")
		fmt.Println("--------------------------------------")
	}
	button := ""
	fmt.Println("Would you like to try again If yes press \"X\" if no press \"N\" ")
	fmt.Fscan(os.Stdin, &button)
	strings.ToLower(button)
	if button == "x" {
		goto startGame
	} else if button == "n" {
		fmt.Println("Bye,Bye.")
	}
}
