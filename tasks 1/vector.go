//Создайте struct который описывает вектор (в трёхмерном пространстве).
// У него должны быть:
//конструктор с параметрами в виде списка координат x, y, z
//метод, вычисляющий длину вектора. Корень можно посчитать с помощью Math.sqrt():
//метод, вычисляющий скалярное произведение:
//метод, вычисляющий векторное произведение с другим вектором:
//метод, вычисляющий угол между векторами (или косинус угла):
// косинус угла между векторами равен скалярному произведению векторов,
//деленному на произведение модулей (длин) векторов
//методы для суммы и разности:
//статический метод, который принимает целое число N,
//и возвращает массив случайных векторов размером N.
//Если метод возвращает вектор, то он должен возвращать новый объект,
//а не менять базовый. То есть, нужно реализовать шаблон

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Vector struct {
	x, y, z float64
}

func newVector(x, y, z float64) Vector { //Конструктор для структуры "Vector"
	return Vector{
		x: x,
		y: y,
		z: z,
	}
}

func (v *Vector) VectorLength() float64 { //метод, вычисляющий длину вектора. Корень
	x := v.x
	y := v.y
	z := v.z
	sum := x*x + y*y + z*z
	return math.Sqrt(sum)
}
func (v Vector) ScalarProduct() float64 { // метод, вычисляющий скалярное произведение
	x := v.x
	y := v.y
	z := v.z
	sum := x*v.x + y*v.y + z*v.z
	return sum
}
func (v Vector) productWithAnotherVector(a Vector) float64 { //векторное произведение с другим вектором
	sum := v.x*a.x + v.y*a.y + v.z*a.z
	return sum
}

func (v Vector) angleBetweenVectors(a Vector) float64 { //метод, вычисляющий угол между векторами (или косинус угла)
	Scalar := v.productWithAnotherVector(a) // скалярное произведение векторов

	aX := math.Sqrt(math.Pow(v.x, 2) + math.Pow(a.x, 2)) //модули векторов
	bY := math.Sqrt(math.Pow(v.y, 2) + math.Pow(a.y, 2)) //модули векторов

	return Scalar / (aX * bY) //math.Cos(Scalar / (aX * bY)) или косинус угла
}
func sliceOfRandomVectors(number int) []Vector { //функция , которыя принимает целое число N,и возвращает массив случайных векторов размером N
	sliceVectors := []Vector{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		num := make([]float64, number)
		for j := 0; j < number; j++ {
			num[j] = math.Round((rand.Float64()*100)*100) / 100
			num = append(num, num[j])
		}
		sliceVectors = append(sliceVectors, newVector(num[0], num[1], num[2]))

	}
	return sliceVectors
}

func main() {

	test := newVector(24, 5, 8)
	test.VectorLength()
	anotherVector := newVector(12, 4, 9)
	test.productWithAnotherVector(anotherVector)
	fmt.Println(test.ScalarProduct(), test.productWithAnotherVector(anotherVector))
	fmt.Println(sliceOfRandomVectors(3))
}
