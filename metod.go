package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Persona struct {
	Id      string
	FIO     string
	Address string
	// Product Product
}

// type Product struct {
// 	Name     string
// 	Quantity int
// 	Price    string
// }
func (p *Persona) addFio(FIO string) {

	p.FIO = FIO

}
func (p *Persona) addAddress(Address string) {

	p.Address = Address

}
func newPersona(Id, FIO, Address string) Persona {
	return Persona{
		Id:      Id,
		FIO:     FIO,
		Address: Address,
	}
}

func main() {
	// open the file
	file, err := os.Open("test/1.txt")

	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	b := [][]byte{}
	for fileScanner.Scan() {
		row := ""
		// fmt.Println(fileScanner.Text())
		row = fileScanner.Text()
		a := []byte(row)

		b = append(b, a)
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

	var persona1, persona2, persona3 Persona
	err = json.Unmarshal(b[0], &persona1)
	err = json.Unmarshal(b[1], &persona2)
	err = json.Unmarshal(b[2], &persona3)

	persona4 := newPersona("", "", "")

	persona4.addFio("Serghei Muntenu")
	persona4.addAddress("Africa")
	fmt.Println(persona1, persona2, persona3, persona4)
}
