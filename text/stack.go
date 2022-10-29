package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hall struct {
	Name     string
	Movie    string
	Discount map[int]int //ключ количество билетов,значение - процент скидки
	Price    map[int]int //ключ - часы, значение - цена
}

type Order struct {
	NumberOfTickets int    // количество заказаных билетов
	Hour            int    // время заказа
	Hall            string // выбранный зал
}

func (o Order) Total() int {
	return 0
}

type Halls []Hall //Все данные хранятся в слайсе
func (h Halls) CheckHall(name string) bool {
	for _, val := range h {
		if val.Name == name {
			return true
		}

	}
	return false
}

func main() {
	var infoOfHalls = Halls{
		{Name: "red hall", Movie: "The Chronicles of Narnia", Discount: map[int]int{}, Price: map[int]int{12: 25, 16: 35, 20: 45}},
		{Name: "Blue Hall", Movie: "Aliens", Discount: map[int]int{}, Price: map[int]int{12: 25, 16: 35, 20: 45}},
		{Name: "Yellow Hall", Movie: "Avatar", Discount: map[int]int{}, Price: map[int]int{12: 25, 16: 35, 20: 45}},
	}

	fmt.Println("Welcome to our cinema")

	for i := 0; i < len(infoOfHalls); i++ {
		var str string

		for time, price := range infoOfHalls[i].Price {

			str += "at " + strconv.Itoa(time) + "o'clock and price " + strconv.Itoa(price) + "$" + "\n"
		}
		fmt.Printf("In _%s_ you can watch the movie\n %3s'%s'\n%s\n", infoOfHalls[i].Name, "", infoOfHalls[i].Movie, str)
	}
	var order Order
	i := 1
	for i > 0 {
		fmt.Println("Enter the hall name:")

		reader := bufio.NewReader(os.Stdin)
		n, _ := reader.ReadString('\n')
		nameHall := strings.TrimSpace(n)

		if !infoOfHalls.CheckHall(nameHall) {
			fmt.Println("You are mistaken")
			i++
		}
		order.Hall = nameHall
	}
}
