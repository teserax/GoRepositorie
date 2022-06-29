//Для каждого билета случайным образом выбирается космическая станция:
//	Space Adventures, SpaceX или Virgin Galactic.
//Датой отправления на каждом билете значится 13 Октября 2020 года.
//В этот день Марс будет на расстоянии 62 100 000 км от Земли.
//Скорость космического корабля будет выбрана случайным
//образом из диапазона от 16 до 30 км/ч. Это определит продолжительность
//поездки на Марс, а также цену билета.
//Более быстрые корабли намного дороже.
//Цены на билеты варьируются от $36 до $50 миллионов.
//Цена для поездки туда-обратно удваивается.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Spaceline struct {
	nameStation string
	days        int
	tripType    string
	price       int
}

const distance = 62100000

func randomNum(start, end int) int {
	num := rand.Intn(start) + end
	return num
}
func duration(speed int) (days int) { //duration days
	return distance / (24 * (3600 * speed))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	station := []string{"Space Adv", "Space XPS", "Space GAL"}
	trip := []string{"Round-trip", "One-way"}
	info := make([]Spaceline, 10) //list of tickets
	for i := 0; i < 10; i++ {     //filling with random data
		info[i].nameStation = station[randomNum(3, 0)]
		info[i].days = (duration((randomNum(15, 15))))
		info[i].tripType = trip[randomNum(2, 0)]
		var extraPrice int
		if info[i].tripType == "Round-trip" {
			extraPrice = 2
		} else {
			extraPrice = 1
		}
		switch {
		case info[i].days <= 30:
			info[i].price = 50 * extraPrice
		case info[i].days > 30 || info[i].days <= 35:
			info[i].price = 43 * extraPrice
		case info[i].days >= 36:
			info[i].price = 36 * extraPrice
		}
	}
	fmt.Println("====================================================")
	fmt.Printf("name station    Days      trip travel     price\n")
	fmt.Println("====================================================")
	for _, ticket := range info {

		fmt.Printf("%s|  %8d| %15s|        $%d \n", ticket.nameStation, ticket.days, ticket.tripType, ticket.price)
	}
	fmt.Println("====================================================")
}
