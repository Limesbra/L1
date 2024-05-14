package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// worker - функция, которая читает данные из канала и выводит их в stdout.
func worker(ch chan string) {
	for {
		msg, ok := <-ch
		// если канал закрыт, то выход из цикла
		if !ok {
			return
		}
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {

	N, err := strconv.Atoi(os.Args[1]) // количество воркеров которое задаем при старте программы
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sig := make(chan os.Signal, 1)     // sig - создание канала для сигналов операционной системы
	signal.Notify(sig, syscall.SIGINT) // signal.Notify - регистрация канала для получения сигнала SIGINT

	data := make(chan string)

	// запуск N потоков worker
	for i := 0; i < N; i++ {
		go worker(data)
	}

	// запуск горутины, которая постоянно записывает в канал данные
	go func() {
		for {
			select {
			// Закрываем канал при получении сигнала
			case <-sig:
				// Закрытие канала данных и завершение функции
				close(data)
				return
			default:
				data <- time.Now().String()
				time.Sleep(time.Second)
			}

		}
	}()

	fmt.Println("\nполученный сигнал -", <-sig, "\nПрограмма завершена") // ожидание сигнала                                                          // отмена операций
}
