package method4

import (
	"fmt"
	"runtime"
	"time"
)

// реализация метода завершение горутины через runtime.Goexit
func Method4() {
	ch := make(chan int)

	// Запуск горутины
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		// Завершение горутины
		runtime.Goexit()

	}()
	// Запуск горутины
	for i := 0; i < 3; i++ {
		go func() {
			<-ch
			fmt.Println("Работаю...")
			// Завершение горутины
			runtime.Goexit()
			for {
				fmt.Println("Работаю... опять")
			}
		}()
	}
	time.Sleep(1 * time.Second)
	close(ch)
	fmt.Println("Остановлен")
	fmt.Println("Method4: Программа завершена")
}
