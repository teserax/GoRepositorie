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

func fullPersona(p *Persona) *Persona {
	infoAddress(p.address.city, p.address.country, p.address.city)
	infoInfo(p.info.fio, p.info.job, p.info.age, p.info.birthday)
	return p
}
func infoInfo(fio, job string, age, birthday int) *Info {
	info := Info{
		fio:      fio,
		job:      job,
		age:      age,
		birthday: birthday,
	}
	return &info
}
func infoAddress(country, city, street string) *Address {
	address := Address{
		country: country,
		city:    city,
		street:  street,
	}
	return &address

}
func main() {

	persona := (&Persona{*infoAddress("st", "sd", "erere"), *infoInfo("fdfd", "", 0, 0)})
	
	fmt.Println(persona)
}
