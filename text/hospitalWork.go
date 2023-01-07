//Структуры данных. Файлы. Дидактический блок 14Основы программирования. Фундаментальный курс | 217
//2. Дана следующая структура:
//No карточки Пациент Заметки
//ФИО Возраст Группа крови Диагноз
//Разработай алгоритм, который:
////2.1. Создает файл с вышеуказанной структурой, информация..............."+"
////о пациенте считывается с клавиатуры.............."+"
//2.2. Выводит содержимое файла на экран................."+"
//2.3. Определяет количество пациентов:,,,,,,,,,,,,,,,"+"
//до 40 лет;........................."+"
//от 41 до 60 лет;..............."+"
//более 60 лет.............."+"
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
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	infoOfPacient := map[int]Patient{}
	start := true
	inc := 0
	data := []byte{}

	for start {

		hospitalQuestionnaire := Patient{}
		fmt.Println("Enter cardNumber of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.CardNumber)
		fmt.Println("Enter FirstName of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.FirstName)
		fmt.Println("Enter LastName of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.LastName)
		fmt.Println("Enter Patronymic of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.Patronymic)
		fmt.Println("Enter Age of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.Age)
		fmt.Println("Enter blood of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.BloodType)
		fmt.Println("Enter diagnosis of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.Diagnosis)

		inc++
		infoOfPacient[inc] = hospitalQuestionnaire
		var answer string

		data, err = json.MarshalIndent(infoOfPacient[inc], " ", "")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Enter New Pacient ? y/n ")
		fmt.Scan(&answer)
		if answer != "y" {
			start = false
		}
	}
	file.Write(data)
	info, err := os.ReadFile("card-index.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(info))
	var fortyYearOlds, fortyToSixty, overSixty int
	for _, v := range infoOfPacient {
		if v.Age <= 40 {
			fortyYearOlds++

		}
		if v.Age > 40 || v.Age <= 60 {
			fortyToSixty++
		}
		if v.Age > 60 {
			overSixty++
		}
		if v.Age <= 50 || v.Diagnosis == "hypertension" {
			fmt.Println(v.FullName, v.Diagnosis)
		}
	}
	fmt.Printf("Number patients under 40 years of age %d \nNumber patients under 40 years of age %d \nNumber patients under 40 years of age %d", fortyYearOlds, fortyToSixty, overSixty)

}
func searchAge(m map[int]Patient) {
	col := 0
	for _, v := range m {
		if v.Age <= 40 {
			col++

		}
	}
	fmt.Printf("Number patients under 40 years of age ", col)
}
