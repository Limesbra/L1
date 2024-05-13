package reverse

import (
	"strings"
)

// Reverse переворачивает строку, разбивая ее на слова и переворачивая их порядок.
// Аргумент s - строка, которую нужно перевернуть.
// Возвращает перевернутую строку.
func Reverse(s string) string {
	// разбиваем строку на слова и сохраняем их в слайс
	strs := strings.Split(s, " ")
	// меняем местами слова в слайсе
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	// объединяем слова в строку
	s = strings.Join(strs, " ")

	return s
}

// Reverse2 переворачивает строку, разбивая ее на слова и переворачивая их порядок.
// Аргумент s - строка, которую нужно перевернуть.
// Возвращает перевернутую строку.
func Reverse2(s string) string {
	var str string
	// находим последний пробел в строке
	spaceIndex := strings.LastIndexByte(s, ' ')
	//пока пробелов в строке не останется
	for spaceIndex != -1 {
		// добавляем слово после пробела в строку
		str = str + s[spaceIndex+1:] + " "
		// удаляем слово после пробела
		s = s[:spaceIndex]
		// находимляем последний пробел в строке
		spaceIndex = strings.LastIndexByte(s, ' ')
	}
	// добавляем остаток строки
	str = str + s
	return str

}
