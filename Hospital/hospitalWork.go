//Структуры данных. Файлы. Дидактический блок 14Основы программирования. Фундаментальный курс | 217
//2. Дана следующая структура:
//No карточки Пациент Заметки
//ФИО Возраст Группа крови Диагноз
//Разработай алгоритм, который:
////2.1. Создает файл с вышеуказанной структурой, информация
////о пациенте считывается с клавиатуры.
//2.2. Выводит содержимое файла на экран.
//2.3. Определяет количество пациентов:
//до 40 лет;
//от 41 до 60 лет;
//более 60 лет.
//2.4. Определяет количество пациентов с гипертонической
//болезнью и возрастом менее 50 лет.
//2.5. Определяет количество пациентов с артритом и группой
//крови A4.
//2.6. Определяет какая группа крови является наиболее рас-
//пространенной при анемии – A2 или A0?
//2.7. Определяет наиболее распространенную группу крови.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Patient struct {
	CardNumber int
	Age        int
	BloodType  string
	Diagnosis  string
	FullName
}
type FullName struct {
	FirstName  string
	LastName   string
	Patronymic string
}

func main() {
	file, err := os.Create("card-index.txt")
	if err != nil {
		log.Fatal("Error make file")
		os.Exit(1)
	}
	defer file.Close()

	infoOfPacient := map[int]Patient{}
	//start := true
	i := 0
	infoOfPacient[i] = Patient{}
	fmt.Println("Hello enter pacient info ")
	fmt.Print("Fist Name") //запись с клавиатуры???
	infoOfPacient[i] = Patient{
		CardNumber: 1,
		Age:        22,
		BloodType:  "AO",
		Diagnosis:  "test",
		FullName: FullName{
			"Ivanov",
			"Serghei",
			"Petrovich",
		},
	}
	data := []byte{}
	data, _ = json.MarshalIndent(infoOfPacient[i], " ", "")
	file.Write(data)
	//if i == 3 {
	//	start = false
	//}

	info, err := os.Open("card-index.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer info.Close()
	dataInfo := make([]byte, 64)

	for {
		n, err := info.Read(dataInfo)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}

}
