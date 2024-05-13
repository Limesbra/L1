package main

import "fmt"

func main() {
	a := 27
	b := 11

	// Способ 1 сложение и вычитание (аналогично умножение и деление)
	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a:", a) // Выведет: a: 11
	fmt.Println("b:", b) // Выведет: b: 27
	fmt.Println()

	//способ 2 оператор присваивания для нескольких переменных
	a, b = b, a

	fmt.Println("a:", a) // Выведет: a: 27
	fmt.Println("b:", b) // Выведет: b: 11
	fmt.Println()

	// способ 3 исключающее или
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("a:", a) // Выведет: a: 11
	fmt.Println("b:", b) // Выведет: b: 27
}
