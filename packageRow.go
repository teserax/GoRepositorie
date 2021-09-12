/*Add package to detect base params of row
добавить пакет определения основных характеристик строки:

количество слов в строке. В качестве разделителя использовать пробел
количество предложений в строке. В качестве разделителя использовать точку, знак восклицания, вопросительный знак.
количество знаков препинания
количество букв.*/
package packagerow

import "strings"

func NumberOfWords(str string) int {
	words := strings.Fields(str)
	var number int
	for i := range words {
		number = i + 1

	}
	return number

}
func SentencesSeparatesCommas(str string) int {
	words := strings.Split(str, ",")
	var number int
	for i := range words {
		number = i

	}
	return number
}
