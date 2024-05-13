package main

import (
	"fmt"
)

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем пустое множество (map) для хранения уникальных строк
	Set := make(map[string]bool)

	// Проходим по каждой строке в исходной последовательности
	for _, str := range strings {
		// Добавляем строку в множество, если она еще не там
		Set[str] = true
	}

	// Выводим уникальные строки из множества
	for str := range Set {
		fmt.Println(str)
	}
}
