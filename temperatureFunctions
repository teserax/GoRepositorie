// // Повторно используйте функцию kelvinToCelsius для конвертации 233° К в градусы Цельсия;
// // Напишите и используйте функцию конвертации температуры в градусы
//  Фаренгейта — celsiusToFahrenheit. Формула для конвертации температуры в градусы по Фаренгейту: (c * 9.0 / 5.0) + 32.0;
// // Напишите функцию kelvinToFahrenheit и проверьте, чтобы она конвертирова 0° К в приблизительно –459.67° F.

package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64 { //функция  для конвертации  Кельвина в градусы Цельсия
	return k - 273.15
}
func celsiusToFahrenheit(c float64) float64 { //Функция для конвертации Цельсия в градусы по Фаренгейту
	return (c * 9.0 / 5.0) + 32.0
}
func kelvinToFahrenheit(k float64) float64 { //Функция для конвертации Цельсия в градусы по Фаренгейту
	if k == 0 {
		k = -459.67
		return k
	}
	return (k-273.15)*9/5 + 32
}

func main() {

	celsius := kelvinToCelsius(294)
	fmt.Printf("K/C %0.2f \n", celsius)
	newCelsius := kelvinToCelsius(233)
	fmt.Printf("K/C %0.2f \n", newCelsius)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("C/F %0.2f \n", fahrenheit)
	fahrenheit = kelvinToFahrenheit(33)
	fmt.Printf("K/F %0.2f \n", fahrenheit)

}
