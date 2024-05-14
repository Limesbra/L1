package main

import (
	"task_6/internal/method1"
	"task_6/internal/method2"
)

func main() {
	method1.Method1()
	method2.Method2()
}

// import ()

// func main() {

// }

// import ()

// func main() {

// }

// import ()

// func main() {

// }

// import ()

// func main() {

// }

// import ()

// func main() {

// }

// Реализовать все возможные способы остановки выполнения горутины.
/*
1) Использование context.Context для отмены
	context.Context предоставляет мощный механизм для управления жизненным циклом горутин и отмены операций.
	Вы можете создать контекст с отменой и передать его в горутину,которая будет проверять, была ли отмена,
	и в зависимости от этого останавливать свою работу.

2) Использование sync.WaitGroup для ожидания завершения
	sync.WaitGroup позволяет ожидать завершения определенного количества горутин. Это не является прямым способом
	остановки горутины, но позволяет корректно завершить программу, ожидая завершения всех горутин.

3) Использование sync.Mutex или sync.RWMutex для блокировки
	Вы можете использовать sync.Mutex или sync.RWMutex для контроля доступа к общим ресурсам и остановки горутины,
	если определенное условие становится истинным.

4) Использование каналов для сигналов
	Каналы могут быть использованы для передачи сигналов между горутинами, включая сигналы для остановки.

5) Использование os.Signal для обработки сигналов от операционной системы
	Вы можете использовать пакет os/signal для обработки сигналов от операционной системы, таких как SIGINT (Ctrl+C), и
	остановить горутину в ответ на эти сигналы.
*/
