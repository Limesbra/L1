package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Создаем необходимые каналы.
	p := producer(arr)
	out := consumer(p)

	// Выводим значения.
	for v := range out {
		fmt.Println(v)
	}

}

// producer - функция, которая возвращает канал с числами
func producer(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// consumer - функция, которая возвращает канал с результатот операции x*2
func consumer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}
