/* Напишите программу, которая делает следующее:

Он запрашивает количество воды, молока и кофейных зерен, доступных в данный момент, а затем запрашивает количество чашек, которое нужно пользователю.
1 Если в кофемашине достаточно расходных материалов для приготовления указанного количества кофе,
    программа должна напечатать "Yes, I can make that amount of coffee".
2 Если кофемашина может приготовить больше, программа должна вывести "Yes, I can make that amount of coffee (and even N more than that)",
    где N — количество дополнительных чашек кофе, которые может приготовить кофемашина.
3 Если количества заданных ресурсов недостаточно для приготовления указанного количества кофе, программа должна вывести "No, I can make only N cups of coffee".
Как и на предыдущем этапе, кофемашине требуется 200 мл воды, 50 мл молока и 15 г кофейных зерен для приготовления одной чашки кофе.*/
package main

import "fmt"

func main() {
	var cups, water, milk, beans, rest int
	fmt.Println("How many ml of water the coffee machine has")
	fmt.Scan(&water)
	fmt.Println("How many ml of milk the coffee machine has:")
	fmt.Scan(&milk)
	fmt.Println("How many grams of coffee beans the coffee machine has:")
	fmt.Scan(&beans)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)
	portionOfWater := water / 200 //находим сколько порции воды
	portionOfMilk := milk / 50    //молока
	portionOfBeans := beans / 15  // кофе
	rest = portionOfWater         // нас интересует минимальная возможная порция
	if rest > portionOfMilk {     //можно в слас и через range ...но мы работаем с логическими операторами
		rest = portionOfMilk
	}
	if rest > portionOfBeans {
		rest = portionOfBeans
	}
	if rest-cups >= 0 { // если общая сумма больше или ровна нулю (минимальная порция минус нужное количество кружек) то проверяем на условие
		if rest-cups > 0 { //если больше то разница составляет сумму возможных кружек
			fmt.Printf("Yes, I can make that amount of coffee,(and even %d more than that)\n", rest-cups) //вот мы и указываем что можно еще сделать
		} else if cups-rest == 0 { //если остаток по-нулям значит мы можем только нужное количество сделать
			fmt.Println("Yes, I can make that amount of coffee")
		}
	} else {
		fmt.Printf("No, I can make only %d cups of coffee\n", rest) //к сожалению мы можем только определенное количество меньше требуемого
	}
}
