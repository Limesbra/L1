package main

import (
	"fmt"
	"sync"
	"task_18/internal/counter"
)

func main() {
	counter := &counter.Counter{}
	finish := make(chan bool)
	go worker(counter, finish)
	<-finish
	fmt.Println(counter.Value())
}

// worker - функция, которая запускает 10 горутин, каждая из которых увеличивает счетчик на единицу.
// После завершения работы всех горутин, функция отправляет сигнал о завершении в канал finish.
func worker(c *counter.Counter, finish chan bool) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	finish <- true
}
