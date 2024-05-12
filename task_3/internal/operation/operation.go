package operation

import "sync"

func square(n int, ch chan<- int) {
	ch <- n * n
}

func SumSqures(numbers []int) {
	ch := make(chan int, 5) // Буферизированный канал емкостью пять элементов
	var wg sync.WaitGroup   // Синхронизация потоков
	wg.Add(5)               // 5 потоков

	for _, number := range numbers { // Запуск горутин
		go square(number, ch)
		wg.Done() // Уменьшение счетчика
	}
	wg.Wait()
	var result int
	for i := 0; i < 5; i++ {
		result += <-ch

	}
	println(result)
}
