package main

import "task_3/internal/operation"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	operation.SumSqures(numbers)
}
