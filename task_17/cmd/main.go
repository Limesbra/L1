package main

import (
	"fmt"
	"task_17/internal/binarysearch"
)

func main() {
	// Создаем отсортированные массивы
	arri := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 90}
	arrs := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "imbe", "jackfruit"}
	arrf := []float32{2.5, 5.2, 8.7, 12.3, 16.8, 23.9, 38.2, 56.4, 72.1, 90.0}

	// Выполняем бинарный поиск и выводим результат в консоль
	place, ok := binarysearch.BinarySearch(arri, 5)
	fmt.Println(place, ok)
	place, ok = binarysearch.BinarySearch(arrs, "nut")
	fmt.Println(place, ok)
	place, ok = binarysearch.BinarySearch(arrf, 2.5)
	fmt.Println(place, ok)

}
