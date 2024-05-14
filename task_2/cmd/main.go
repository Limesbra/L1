package main

import (
	"fmt"
	"sort"
	"sync"

	"task_2/internal/square"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 5)
	var wg sync.WaitGroup // Синхронизация потоков
	var mu sync.Mutex     // Блокировка доступа к общим ресурсам
	wg.Add(5)             // 5 потоков

	for i := 0; i < 5; i++ { // Запуск горутин
		go square.Square(numbers[i], ch, &mu)
		wg.Done() // Уменьшение счетчика
	}
	wg.Wait() // Ожидание завершения работы потоков

	res := make([]int, 5)

	for i := 0; i < 5; i++ {
		res[i] = <-ch
	}

	sort.Ints(res)

	for _, num := range res { // Вывод результата
		fmt.Printf("%d ", num)
	}
}
