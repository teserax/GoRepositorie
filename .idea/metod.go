
//Есть например файл такого содержания. Это ID и имя человека
//```
//{"id":"1", "FIO":"Galits Valera"}
//{"id":"2", "FIO":"Ciapa Vladimir"}
//{"id":"3", "FIO":"Ciapa Rocsana"}
//
//```
//
//второй файл с адресом человека
//
//
//{"id":"1", "address":"Baltsy"}
//{"id":"2", "address":"Chisinau"}
//{"id":"3", "address":"Chisinau"}
//
//
//третий с покупками человвека
//
//
//{"id":"1","product":{"id","Vodka","quantity":10,"price":"5"}}
//{"id":"2","product":{"id","Vodka","quantity":10,"price":"10"}}
//{"id":"1","product":{"id","Vodka","quantity":10,"price":"10"}}
//{"id":"2","product":{"id","Vodka","quantity":10,"price":"10"}}
//{"id":"3","product":{"id","Beer","quantity":10,"price":"10"}}
//{"id":"3","product":{"id","Beer","quantity":10,"price":"10"}}
//{"id":"3","product":{"id","Beer","quantity":10,"price":"10"}}
//{"id":"3","product":{"id","Beer","quantity":10,"price":"10"}}
//{"id":"3","product":{"id","Beer","quantity":10,"price":"10"}}`
//
//Нужно получить объединяющий файл куда записать объединенные данные о человеке - ID,Name,Address, Products (товары которые он купил). Заполнения финальной структуры сделать через методы - метод заполняющий ID  и Name, второй метод для Адреса, третий для товара.
//
//то есть будет примерно так
//
//person := New()
//
//person.AddFIO(...)
//person.AddAddress(...)
//person.AddProduct(....)
package metod

import "fmt"

type Person struct{
	Id string
	FIO string
	address string
	product Product

}
type Product struct {
	productNane string
	quantity int
	price string
}

func main(){

	buyer:= Person{"1","Galit Valera","balti",Product{"beer",10,"10"}}
	fmt.Println(buyer)
}
