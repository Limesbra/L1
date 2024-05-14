package method2

import (
	"context"
	"fmt"
	"time"
)

func Method2() {
	// Создание контекста для передачи сигналов об остановке горутины
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				// Получен сигнал об остановке
				fmt.Println("Остановлен")
				return
			default:
				// Выполнение работы в горутине
				fmt.Println("Работаю...")
				time.Sleep(1 * time.Second)
			}
		}

	}()

	time.Sleep(3 * time.Second)
	// Отмена контекста
	cancel()
	// Ждем завершения горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Method2: Программа завершена")
}
