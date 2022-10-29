//Cоздать программу (простую) в которой будешь заполнять данными структурус 2мя полями:
//	address (country, city, street), info (fio, job, age, birthday).
//	Поле заполнение должно быть через 3 функции:
//	первая заполняет всю структуру,
//    внутри вызывает 2 функции - одна заполняет address, вторая поле info.

package main

import "fmt"

type Persona struct {
	address Address
	info    Info
}
type Address struct {
	country string
	city    string
	street  string
}
type Info struct {
	fio      string
	job      string
	age      int
	birthday int
}

func fullPersona(p *Persona) {
	infoAddress(&p.address)
	infoInfo(&p.info)
}
func infoInfo(inf *Info) {
	inf.fio = "Peter"
	inf.job = "Prezident"
	inf.age = 78
	inf.birthday = 2022
}
func infoAddress(addr *Address) {
	addr.city = "Chisinau"
	addr.country = "Moldova"
	addr.street = "durban"
}
func main() {
	var pers Persona

	fullPersona(&pers)
	infoInfo(&pers.info)
	infoAddress(&pers.address)
	fmt.Println(pers)
}
