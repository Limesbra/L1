package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	m := make(map[int]int)

	// Цикл для заполнения карты значениями от 1 до 10
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// Запуск горутины для обновления значения в карте
		go func(i int) {
			// блокировка мьютекса
			mu.Lock()
			m[i] = i + 1
			wg.Done()
			// разблокировка мьютекса
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(m)
}
