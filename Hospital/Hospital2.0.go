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
type List []Patient
type Options struct {
	MinAge    int
	MaxAge    int
	BloodType string
	Diagnosis string
}

func (list List) Filter(option Options) List {

	var result List
	for _, patient := range list {
		if (patient.Age >= option.MinAge || option.MinAge == 0) &&
			(patient.Age <= option.MaxAge || option.MaxAge == 0) &&
			(patient.BloodType == option.BloodType || option.BloodType == "") &&
			(patient.Diagnosis == option.Diagnosis || option.Diagnosis == "") {

			result = append(result, patient)
		}
	}
	return result

}
func (list List) widespread(option Options) {
	widesp := map[string]int{}
	for _, patient := range list {
		if option.BloodType == "" {
			if _, exist := widesp[patient.BloodType]; exist {
				widesp[patient.BloodType]++
			} else {
				widesp[patient.BloodType] = 1
			}
		}
		if option.Diagnosis == "" {
			if _, exist := widesp[patient.Diagnosis]; exist {
				widesp[patient.Diagnosis]++
			} else {
				widesp[patient.Diagnosis] = 1
			}
		}

	}
	max := 0
	value := ""
	for i, val := range widesp {
		if val > max {
			max = val
			value = i
		}
	}
	fmt.Printf("amount %d value %s\n", max, value)

}

func main() {
	file, err := os.Create("card-index.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	ListOfPacient := List{}
	//start := true
	data := []byte{}

	//for start {

	//hospitalQuestionnaire := Patient{}
	//fmt.Println("Enter cardNumber of Pacient ")
	//fmt.Scan(&hospitalQuestionnaire.CardNumber)
	//fmt.Println("Enter FirstName of Pacient ")
	//fmt.Scan(&hospitalQuestionnaire.FirstName)
	//fmt.Println("Enter LastName of Pacient ")
	//fmt.Scan(&hospitalQuestionnaire.LastName)
	//fmt.Println("Enter Patronymic of Pacient ")
	//fmt.Scan(&hospitalQuestionnaire.Patronymic)
	//fmt.Println("Enter Age of Pacient ")
	//fmt.Scan(&hospitalQuestionnaire.Age)
	//fmt.Println("Enter blood of Pacient ")
	//fmt.Scan(&hospitalQuestionnaire.BloodType)
	//fmt.Println("Enter diagnosis of Pacient ")
	//fmt.Scan(&hospitalQuestionnaire.Diagnosis)
	//
	//ListOfPacient = append(ListOfPacient, hospitalQuestionnaire)
	//
	ListOfPacient = []Patient{
		{1, 15, "AO", "test", FullName{FirstName: "Petrov", LastName: "Serghei", Patronymic: "qww"}},
		{2, 32, "AA", "test", FullName{FirstName: "Ivanov", LastName: "Andrei", Patronymic: "qggg"}},
		{3, 56, "AO", "top", FullName{FirstName: "Sidorov", LastName: "Petar", Patronymic: "qhhh"}},
		{4, 67, "AB", "test", FullName{FirstName: "Kozlov", LastName: "Ivan", Patronymic: "qkkk"}},
	}

	data, err = json.MarshalIndent(ListOfPacient, " ", "")
	if err != nil {
		log.Fatal(err)
	}
	//	var answer string
	//	fmt.Println("Enter New Pacient ? y/n ")
	//	fmt.Scan(&answer)
	//	if answer != "y" {
	//		start = false
	//	}
	//}
	file.Write(data)
	info, err := os.ReadFile("card-index.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	fmt.Println(string(info))
	/*infoResult:= ListOfPacient.Filter(Options{MaxAge: 40})
	fmt.Printf("found  patients up to 40 years of age: %v\n",  infoResult)
	*/
	infoResult := ListOfPacient.Filter(Options{MinAge: 11, MaxAge: 60})
	fmt.Printf("found %d patients from 41 up to 60 years of age: %v\n", len(infoResult), infoResult)
	/*
		infoResult= ListOfPacient.Filter(Options{MinAge: 60, MaxAge: 160})
		fmt.Printf("found patients  up to 60 years of age: %v\n",  infoResult)*/
	//	infoResult := ListOfPacient.Filter(Options{Diagnosis: "test"}).Filter(Options{BloodType: "00"})
	//	fmt.Printf("found  patients with this diagnosis \n%v\n", infoResult)
	ListOfPacient.widespread(Options{BloodType: ""})
	//ListOfPacient.widespread(Options{Diagnosis: ""})

}
