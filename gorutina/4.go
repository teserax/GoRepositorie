package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var winner = make(chan string) // инициализация канала winner
var echo = make(chan struct{}) //инициализация канала echo

func game_table(s string) {
	r := rand.Intn(10) + 1 //  получение рандомного числа для временнова отрезка
	fmt.Println(s, r)      // имя и полученное число
	//t := time.NewTimer(time.Duration(r) * time.Microsecond) //временной отрезок

	select { // выбор условии
	case winner <- s:
		close(echo)
		fmt.Println(s, "press botton") // имя перврго кто нажал
		//defer t.Stop()
		defer close(winner)

		wg.Done()

		// останавливаем дефером  таймер для имени(вошедшего первым)

	case <-echo: // второй кейс если канал echo  закрыт то
		fmt.Println("Закрыто для ", s) // обьявляем для кого закрылся канал
		//defer t.Stop()
		wg.Done() //останавливаем дефером  таймер для имени(вошедшего )

	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(3)
	go game_table("Max")    //запуск горутины Макс
	go game_table("Vova")   //запуск горутины Вова
	go game_table("Valera") //запуск горутины Валера

	s := <-winner                    //присваеваем имя победителя полученное из канала
	fmt.Println("The winner is ", s) //обьевляем победителя
	wg.Wait()
	//<-time.NewTicker(time.Duration(10) * time.Millisecond).C
}
