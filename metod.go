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
	Product Product
}

type Product struct {
	Goods    string
	Quantity int
	Price    string
}

func (p *Persona) addFio(FIO string) {

	p.FIO = FIO

}
func (p *Persona) addAddress(Address string) {

	p.Address = Address

}
func newPersona(Id, FIO, Address string, Goods string, Quantity int, Price string) Persona {
	return Persona{
		Id:      Id,
		FIO:     FIO,
		Address: Address,
		Product: Product{
			Goods:    Goods,
			Quantity: Quantity,
			Price:    Price},
	}
}

func main() {
	file1, err := readFile("text/1.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	var persona1, persona2, persona3 Persona
	err = json.Unmarshal(file1[0], &persona1)
	err = json.Unmarshal(file1[1], &persona2)
	err = json.Unmarshal(file1[2], &persona3)
	file2, err := readFile("text/2.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	err = json.Unmarshal(file2[0], &persona1)
	err = json.Unmarshal(file2[1], &persona2)
	err = json.Unmarshal(file2[2], &persona3)
	file3, err := readFile("text/3.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	err = json.Unmarshal(file3[0], &persona1.Product)
	err = json.Unmarshal(file3[1], &persona2.Product)
	err = json.Unmarshal(file3[2], &persona3.Product)
	fmt.Println(persona1)

	persona4 := newPersona("4", "", "", "", 0, "")

	persona4.addFio("Serghei Muntenu")
	persona4.addAddress("Africa")
	fmt.Println(persona1, persona2, persona3, persona4)
}

func readFile(a string) ([][]byte, error) {
	// open the file
	file, err := os.Open(a)

	if err != nil {
		log.Fatalf("Error %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	bd := [][]byte{}
	for fileScanner.Scan() {
		row := ""
		// fmt.Println(fileScanner.Text())
		row = fileScanner.Text()
		a := []byte(row)

		bd = append(bd, a)
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	file.Close()
	return bd, nil
}
