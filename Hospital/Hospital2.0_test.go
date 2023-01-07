package main

import (
	"testing"
)

func TestList_Filter(t *testing.T) {
	ListOfPacient := List{}
	ListOfPacient = []Patient{
		{1, 15, "333", "test", FullName{FirstName: "Petrov", LastName: "Serghei", Patronymic: "qww"}},
		{2, 32, "qq", "est", FullName{FirstName: "Ivanov", LastName: "Andrei", Patronymic: "qggg"}},
		{3, 56, "ee", "top", FullName{FirstName: "Sidorov", LastName: "Petar", Patronymic: "qhhh"}},
		{4, 67, "00", "test", FullName{FirstName: "Kozlov", LastName: "Ivan", Patronymic: "qkkk"}},
	}
	tests := []struct {
		search Options
		exp    int
	}{
		{Options{MinAge: 17}, 3},
		{Options{MaxAge: 67}, 4},
		{Options{BloodType: "00", Diagnosis: "test"}, 1},
		{Options{}, 4},
		{Options{Diagnosis: ""}, 0},
	}

	for _, e := range tests {
		res := ListOfPacient.Filter(e.search)
		if len(res) != e.exp {
			t.Errorf("Find( %d), expected %d", len(res), e.exp)
		}
	}
}
