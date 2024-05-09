package main

import (
	"fmt"
	"sort"
	"sync"

	"task_2/internal/square"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 5) // Буферизированный канал емкостью пять элементов
	var wg sync.WaitGroup   // Синхронизация потоков
	wg.Add(5)               // 5 потоков

	for _, number := range numbers { // Запуск горутин
		go square.Square(number, ch)
		wg.Done() // Уменьшение счетчика
	}
	wg.Wait() // Ожидание завершения работы потоков

	result := make([]int, 5)
	for i := 0; i < 5; i++ { // Чтение данных из канала
		result[i] = <-ch
	}

	sort.Ints(result) // Сортировка результата

	for _, num := range result { // Вывод результата
		fmt.Printf("%d ", num)
	}

}
