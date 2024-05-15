package main

// Проблема в данной задачи заключалась в том, что во время выполнения программы
// в момент выполнения функции someFunc() создается большая строка
// и далее присваивается подстрока из большой строки в переменную justString
// Большая строка будет храниться в памяти до завершения программы,
// поскольку justString все еще ссылается на эту большую строку.

import "fmt"

var justString string

func createHugeString() string {
	var s string
	for i := 0; i < 1024; i++ {
		s += "a"
	}
	return s
}

func someFunc() {
	v := createHugeString()
	justStringBytes := make([]byte, 100)

	// чтобы justString не ссылалась на большую строку,
	//с помощью copy создаем новую строку, которая содержит
	//только первые 100 элементов
	copy(justStringBytes, v[:100])
	justString = string(justStringBytes)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
