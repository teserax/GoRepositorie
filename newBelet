package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hall struct {
	Movie    string
	Discount map[int]int //ключ количество билетов,значение - процент скидки
	Price    map[int]int //ключ - часы, значение - цена
}

type Order struct {
	NumberOfTickets int    // количество заказаных билетов
	Hour            int    // время заказа
	Hall            string // выбранный зал
	Movie           string //выбраный фильм
	cost            int
}

func (o *Order) Total() int {
	var total, percent int

	if o.NumberOfTickets >= 5 {
		percent = o.cost * o.NumberOfTickets / 20
		total = o.cost*o.NumberOfTickets - percent
		return total
	} else if o.NumberOfTickets >= 10 {
		percent = o.cost * o.NumberOfTickets / 10
		total = o.cost*o.NumberOfTickets - percent
	}
	total = o.cost * o.NumberOfTickets
	return total

}

type Halls map[string]Hall //Все данные хранятся в слайсе
func infoTimePrice(h Hall) string {
	var str string
	for time, price := range h.Price {

		str += "at " + strconv.Itoa(time) + " o'clock and price " + strconv.Itoa(price) + "$" + "\n"
	}
	return str
}
func (h Halls) CheckHall(name string) bool { //Проверка наименования зала
	for n := range h {
		if n == name {
			return true
		}

	}
	return false
}

func main() {
	var infoOfHalls = Halls{
		"red hall":    {Movie: "The Chronicles of Narnia", Discount: map[int]int{}, Price: map[int]int{12: 25, 16: 35, 20: 45}},
		"blue hall":   {Movie: "Aliens", Discount: map[int]int{}, Price: map[int]int{13: 25, 17: 35, 21: 45}},
		"yellow hall": {Movie: "Avatar", Discount: map[int]int{}, Price: map[int]int{10: 25, 14: 35, 19: 45}},
	}

	fmt.Println("Welcome to our cinema")
	var order Order
	for name := range infoOfHalls {
		var str string
		str += infoTimePrice(infoOfHalls[name])

		fmt.Printf("%5sIn _%s_ \nyou can watch the movie\n %3s'%s'\n%s\n", "", name, "", infoOfHalls[name].Movie, str)
	}

	exit := true
	for exit { //цикл запроса выбора зала

		fmt.Println("Enter the hall name:")

		reader := bufio.NewReader(os.Stdin)
		n, _ := reader.ReadString('\n')
		nameHall := strings.TrimSpace(n)

		if !infoOfHalls.CheckHall(nameHall) {
			fmt.Println("You are mistaken")
			continue
		}
		order.Hall = nameHall
		order.Movie = infoOfHalls[nameHall].Movie
		var answer string //запрос на подтвержение о выборе зала
		fmt.Printf("You choose %s and movie --%s-- rightn? Y/N\n", order.Hall, order.Movie)
		fmt.Scanln(&answer)
		if answer != "y" {
			continue

		}
		fmt.Println(infoTimePrice(infoOfHalls[nameHall]))
		exitSessionTime := true
		for exitSessionTime {
			fmt.Println("Please enter a session time")
			fmt.Scanln(&order.Hour)

			for timeExist := range infoOfHalls[order.Hall].Price {
				if timeExist == order.Hour {
					exitSessionTime = false
					break
				} else {
					exitSessionTime = true
				}
			} ///разобраться почему нет совпадения в цикле

			order.cost = infoOfHalls[order.Hall].Price[order.Hour]
			if exitSessionTime == false {
				answer = ""
				fmt.Printf("You choose %d rightn? Y/N\n", order.Hour)
				fmt.Scanln(&answer)
				if answer != "y" {
					exitSessionTime = true
				}
			}
		}

		for order.NumberOfTickets <= 0 {
			fmt.Println("Please enter the number of tickets")
			fmt.Scanln(&order.NumberOfTickets)
		}
		fmt.Printf("You choose %s and movie %s \n ", order.Hall, order.Movie)
		fmt.Printf("Your amount due %d\n", order.Total())
		var finalAnswer string
		fmt.Println("Would you like to choose another room?  Y/N ")
		fmt.Scanln(&finalAnswer)
		if finalAnswer == "y" {
			continue
		} else {
			fmt.Println("Thank you for visiting our cinema. See you again. Goodbye")
			break
		}
	}
}
