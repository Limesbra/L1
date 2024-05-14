package main

import (
	"fmt"
	"sync"
	"task_3/internal/operation"
	"time"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	var (
		mu     sync.Mutex
		result int
	)
	ch := make(chan int, 5)
	numbers := []int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++ {
		go operation.SumSqures(&mu, numbers[i], ch)
	}
	for i := 0; i < 5; i++ {
		go operation.AddToResult(&result, ch, &mu)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(result)
}
