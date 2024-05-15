package method3

import (
	"fmt"
	"sync"
)

func Method3() {
	// Создание WaitGroup для ожидания завершения работы горутин
	wg := sync.WaitGroup{}

	// Добавляем кол-во горутин
	wg.Add(1)

	go func() {
		// Уменьшаем кол-во горутин
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("работаю...")
		}
		fmt.Println("работаю...")
	}()
	// Ожидаем завершения работы горутин
	wg.Wait()
	fmt.Println("Остановлен")
	fmt.Println("Method3: Программа завершена")
}
