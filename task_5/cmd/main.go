package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	duration := time.Duration(N) * time.Second // duration - время работы программы
	timer := time.After(duration)              // timer - канал с таймером

	ch := make(chan string) // канал для передачи данных
	done := make(chan bool) // канал для завершения работы программы

	// запуск горутины, которая читает данные из канала и выводит их в stdout
	go func() {
		for {
			msg, ok := <-ch
			// если канал закрыт, то выход из цикла
			if !ok {
				return
			}
			fmt.Println(msg)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// запуск горутины, которая постоянно записывает в канал данные
	go func() {
		for {
			select {
			// После получения сигнала от таймера закрывает канал и выходит из цикла
			case <-timer:
				close(ch)
				done <- true
				return
			default:
				ch <- time.Now().String()
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	<-done
}
