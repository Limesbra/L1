package main

import (
	"fmt"
	"task_16/internal/quicksort"
)

func main() {
	// Создаем неотсортированные массивы
	arri := []int{5, 2, 8, 1, 9}
	arrf := []float32{5.1, 22.9, -8.9999, 1.000787, -9.0}
	arrs := []string{"apple", "banana", "watermelon", "date", "grape", "pineapple"}

	// Сортируем массивы с помощью функции быстрой сортировки
	sortedArri := quicksort.QuickSort(arri, 0)
	sortedArrf := quicksort.QuickSort(arrf, 4)
	sortedArrs := quicksort.QuickSort(arrs, 2)

	// Выводим отсортированные массивы
	fmt.Println("Отсортированный массив:")
	fmt.Println(sortedArri)
	fmt.Println(sortedArrf)
	fmt.Println(sortedArrs)
}
