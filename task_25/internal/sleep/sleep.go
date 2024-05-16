package sleep

import (
	"time"
)

// Sleep - функция sleep
// реализация через таймер time.After
func Sleep(d time.Duration) {
	<-time.After(d)
}

// Sleep2 - функция sleep
// реализация через цикл
func Sleep2(d time.Duration) {
	// получаем текущее время
	timeStart := time.Now()
	for {
		// пока разница между текущим временем и timeStart меньше d
		// выполняем цикл
		if time.Since(timeStart) >= d {
			return
		}
	}
}
