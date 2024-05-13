package reverse

// Reverse - функция переворачивает строку
// s - строка для переворачивания
// return - перевернутая строка
func Reverse(s string) string {
	runes := []rune(s)                                    // преобразование строки в слайс рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // переворачивание слайса
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) // преобразование слайса рун в строку
}

// Reverse2 - функция переворачивает строку
// s - строка для переворачивания
// return - перевернутая строка
func Reverse2(s string) string {
	output := ""
	for _, item := range s { // переворачивание слайса рун в строку
		output = string(item) + output // добавление элемента вначало строки
	}
	return output

}
